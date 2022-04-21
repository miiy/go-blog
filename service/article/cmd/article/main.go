package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	articlepb "goblog.com/api/article/v1"
	"goblog.com/pkg/grpc_gateway/gateway"
	"google.golang.org/protobuf/encoding/protojson"

	//grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	articleservice "goblog.com/service/article/internal/service"
	//"goblog.com/pkg/jwtauth"
	"goblog.com/service/article/openapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

func main() {
	// flag
	conf := flag.String("c", "./configs/default.yaml", "config file")
	flag.Parse()

	app, cleanUp, err := InitApplication(*conf)
	if err != nil {
		panic(err)
	}
	defer cleanUp()

	zapLogger := app.Logger
	httpAddr := ":50052"
	grpcAddr := ":8082"


	// 不使用 grpcLog，不替换
	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	//grpc_zap.ReplaceGrpcLoggerV2(zapLogger)



	// gRPC server
	lis, err := net.Listen("tcp", httpAddr)
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
			grpcctxtags.StreamServerInterceptor(),
			grpczap.StreamServerInterceptor(zapLogger),
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
			grpcctxtags.UnaryServerInterceptor(),
			grpczap.UnaryServerInterceptor(zapLogger),
			grpc_validator.UnaryServerInterceptor(),
		)),
	)

	articlepb.RegisterArticleServiceServer(s, articleservice.NewArticleServiceServer(app.Database.Gorm, app.Logger))

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
		Addr:                    httpAddr,
		GRPCServerAddr:          grpcAddr,
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
