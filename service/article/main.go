package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"

	//grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"goblog.com/pkg/config"
	//"goblog.com/pkg/jwtauth"
	pkg_zap "goblog.com/pkg/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"goblog.com/pkg/database"
	v1Article "goblog.com/service/article/proto/v1"
	v1ArticleSvr "goblog.com/service/article/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// flag
	configFile := flag.String("f", "../../config/default.yaml", "config file")
	var addr = flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	// config
	c, err := config.NewConfig(*configFile)
	if err != nil {
		log.Fatalln(err)
	}

	zapLogger, err := pkg_zap.NewZap()
	if err != nil {
		log.Fatal(err)
	}
	// 不使用 grpcLog，不替换
	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	//grpc_zap.ReplaceGrpcLoggerV2(zapLogger)

	// db
	db, err := database.NewDatabase(&c.Database)
	if err != nil {
		log.Fatalln(err)
	}

	// gRPC server
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	fmt.Println("Server listening at ", lis.Addr())

	// recovery
	var recoveryHandler grpc_recovery.RecoveryHandlerFunc
	recoveryHandler = func(p interface{}) (err error) {
		zapLogger.Error(fmt.Sprint(p))
		return status.Errorf(codes.Internal, "internal error")
	}
	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryHandler),
	}


	//// jwt
	//jwt := jwtauth.NewJWTAuth(&jwtauth.Options{
	//	Secret:    c.Jwt.Secret,
	//	ExpiresIn: c.Jwt.ExpiresIn,
	//})
	//
	//authRepo := repository.NewRepository(db.Gorm)

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(recoveryOpts...),
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zapLogger),
			grpc_validator.StreamServerInterceptor(),
			//grpc_auth.StreamServerInterceptor(jwt.GrpcAuthenticateInterceptor(authRepo)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_validator.UnaryServerInterceptor(),
			//grpc_auth.UnaryServerInterceptor(jwt.GrpcAuthenticateInterceptor(authRepo)),
		)),
	)

	v1Article.RegisterArticleServiceServer(s, v1ArticleSvr.NewArticleServiceServer(db.Gorm, zapLogger))

	// run GRPC server
	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// run HTTP gateway
	err = RunGateway("localhost:50051", "127.0.0.1:8051")
	if err != nil {
		log.Fatalln(err)
	}
}

func RunGateway(dialAddr, gatewayAddr string) error {
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(context.Background(), dialAddr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	// https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_your_gateway/
	gwMux := runtime.NewServeMux(
		//runtime.WithForwardResponseOption(allowCORS),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions:   protojson.MarshalOptions{
				UseProtoNames: true, // 使用proto中的大小写
				EmitUnpopulated: true, // 显示未填充的字段
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	err = v1Article.RegisterArticleServiceHandler(context.Background(), gwMux, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	srv := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				gwMux.ServeHTTP(w, r)
				return
			}
			// open api
		}),
	}

	log.Printf("Serving gRPC-Gateway on http://%s\n", gatewayAddr)
	return srv.ListenAndServe()
}


