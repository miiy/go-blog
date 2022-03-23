package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"goblog.com/pkg/grpc_gateway/gateway"
	"google.golang.org/protobuf/encoding/protojson"
	//grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"goblog.com/pkg/config"
	"goblog.com/pkg/database"
	//"goblog.com/pkg/jwtauth"
	pkg_zap "goblog.com/pkg/zap"
	"goblog.com/service/article/openapi"
	articlepb "goblog.com/service/article/proto/v1"
	v1ArticleSvr "goblog.com/service/article/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(recoveryOpts...),
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zapLogger),
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_validator.UnaryServerInterceptor(),
		)),
	)

	articlepb.RegisterArticleServiceServer(s, v1ArticleSvr.NewArticleServiceServer(db.Gorm, zapLogger))

	// run GRPC server
	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	serverMuxOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions:   protojson.MarshalOptions{
			UseProtoNames: true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	gwOpts := gateway.Options{
		Addr:                    "127.0.0.1:8051",
		GRPCServerAddr:          "localhost:50051",
		OpenAPIFS:               openapi.OpenAPIFS,
		ServerMuxOption:         []runtime.ServeMuxOption{serverMuxOption},
		RegisterServiceHandlers: []gateway.RegisterServiceHandler{articlepb.RegisterArticleServiceHandler},
	}
	// run HTTP gateway
	err = gateway.Run(context.Background(), gwOpts)
	if err != nil {
		log.Fatalln(err)
	}
}
