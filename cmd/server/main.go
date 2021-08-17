package main

import (
	"flag"
	"fmt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/miiy/go-blog/pkg/api/auth/repository"
	v1Example "github.com/miiy/go-blog/pkg/api/example/proto/v1"
	v1ExampleSrv "github.com/miiy/go-blog/pkg/api/example/service/v1"
	"github.com/miiy/go-blog/pkg/config"
	"github.com/miiy/go-blog/pkg/jwtauth"
	"github.com/miiy/go-blog/pkg/protocol/gateway"
	zap2 "github.com/miiy/go-blog/pkg/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	auth "github.com/miiy/go-blog/pkg/api/auth/proto"
	authSrv "github.com/miiy/go-blog/pkg/api/auth/service"
	feedback "github.com/miiy/go-blog/pkg/api/feedback/proto"
	feedbackSrv "github.com/miiy/go-blog/pkg/api/feedback/service"
	v1Tag "github.com/miiy/go-blog/pkg/api/tag/proto/v1"
	v1TagSrv "github.com/miiy/go-blog/pkg/api/tag/service/v1"
	userPost "github.com/miiy/go-blog/pkg/api/userpost/proto"
	userPostSrv "github.com/miiy/go-blog/pkg/api/userpost/service"
	userTag "github.com/miiy/go-blog/pkg/api/usertag/proto"
	userTagSrv "github.com/miiy/go-blog/pkg/api/usertag/service"
	"github.com/miiy/go-blog/pkg/database"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// flag
	configFile := flag.String("-f", "./configs/default.yaml", "config file")
	var addr = flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	// config
	c, err := config.NewConfig(*configFile)
	if err != nil {
		log.Fatalln(err)
	}

	zapLogger, err := zap2.NewZap()
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






	// jwt
	jwt := jwtauth.NewJWTAuth(&jwtauth.Options{
		Secret:    c.Jwt.Secret,
		ExpiresIn: c.Jwt.ExpiresIn,
	})

	authRepo := repository.NewRepository(db.DB)

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(recoveryOpts...),
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zapLogger),
			grpc_validator.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(jwt.GrpcAuthenticateInterceptor(authRepo)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(jwt.GrpcAuthenticateInterceptor(authRepo)),
		)),
	)

	v1Example.RegisterExampleServiceServer(s, v1ExampleSrv.NewExampleServiceServer())

	// SignUp Tag on the server.
	v1tagOpts := &v1TagSrv.Options{
		Debug: false,
	}
	// SignUp xx on the same server.
	v1Tag.RegisterTagServiceServer(s, v1TagSrv.NewTagServiceServer(v1tagOpts, db.DB))


	auth.RegisterAuthServiceServer(s, authSrv.NewAuthServiceServer(db.DB, jwt))
	feedback.RegisterFeedbackServiceServer(s, feedbackSrv.NewFeedbackServiceServer(db.DB))
	userTag.RegisterUserTagServiceServer(s, userTagSrv.NewUserTagServiceServer(db.DB))
	userPost.RegisterUserPostServiceServer(s, userPostSrv.NewUserPostServiceServer(db.DB))

	// run GRPC server
	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// run HTTP gateway
	err = gateway.Run("localhost:50051", "127.0.0.1:8051")
	if err != nil {
		log.Fatalln(err)
	}
}


