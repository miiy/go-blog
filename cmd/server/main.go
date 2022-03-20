//package main
//
//import (
//	"flag"
//	"fmt"
//	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
//	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
//	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
//	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
//	"goblog.com/pkg/config"
//	"goblog.com/pkg/jwtauth"
//	"goblog.com/pkg/protocol/gateway"
//	zap2 "goblog.com/pkg/zap"
//	"goblog.com/service/auth/repository"
//	v1Example "goblog.com/service/example/proto/v1"
//	v1ExampleSrv "goblog.com/service/example/service/v1"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//
//	"github.com/grpc-ecosystem/go-grpc-middleware"
//	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
//	"goblog.com/pkg/database"
//	auth "goblog.com/service/auth/proto"
//	authSrv "goblog.com/service/auth/service"
//	feedback "goblog.com/service/feedback/proto"
//	feedbackSrv "goblog.com/service/feedback/service"
//	v1Tag "goblog.com/service/tag/proto/v1"
//	v1TagSrv "goblog.com/service/tag/service/v1"
//	userPost "goblog.com/service/userpost/proto"
//	userPostSrv "goblog.com/service/userpost/service"
//	userTag "goblog.com/service/usertag/proto"
//	userTagSrv "goblog.com/service/usertag/service"
//	"google.golang.org/grpc"
//	"log"
//	"net"
//)
//
//func main() {
//	// flag
//	configFile := flag.String("-f", "./config/default.yaml", "config file")
//	var addr = flag.String("addr", "localhost:50051", "the address to connect to")
//	flag.Parse()
//
//	// config
//	c, err := config.NewConfig(*configFile)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	zapLogger, err := zap2.NewZap()
//	if err != nil {
//		log.Fatal(err)
//	}
//	// 不使用 grpcLog，不替换
//	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
//	//grpc_zap.ReplaceGrpcLoggerV2(zapLogger)
//
//	// db
//	db, err := database.NewDatabase(&c.Database)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	// gRPC server
//	lis, err := net.Listen("tcp", *addr)
//	if err != nil {
//		log.Fatalln("Failed to listen:", err)
//	}
//	fmt.Println("Server listening at ", lis.Addr())
//
//	// recovery
//	var recoveryHandler grpc_recovery.RecoveryHandlerFunc
//	recoveryHandler = func(p interface{}) (err error) {
//		zapLogger.Error(fmt.Sprint(p))
//		return status.Errorf(codes.Internal, "internal error")
//	}
//	recoveryOpts := []grpc_recovery.Option{
//		grpc_recovery.WithRecoveryHandler(recoveryHandler),
//	}
//
//
//
//
//
//
//	// jwt
//	jwt := jwtauth.NewJWTAuth(&jwtauth.Options{
//		Secret:    c.Jwt.Secret,
//		ExpiresIn: c.Jwt.ExpiresIn,
//	})
//
//	authRepo := repository.NewRepository(db.DB)
//
//	s := grpc.NewServer(
//		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
//			grpc_recovery.StreamServerInterceptor(recoveryOpts...),
//			grpc_ctxtags.StreamServerInterceptor(),
//			grpc_zap.StreamServerInterceptor(zapLogger),
//			grpc_validator.StreamServerInterceptor(),
//			grpc_auth.StreamServerInterceptor(jwt.GrpcAuthenticateInterceptor(authRepo)),
//		)),
//		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
//			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
//			grpc_ctxtags.UnaryServerInterceptor(),
//			grpc_zap.UnaryServerInterceptor(zapLogger),
//			grpc_validator.UnaryServerInterceptor(),
//			grpc_auth.UnaryServerInterceptor(jwt.GrpcAuthenticateInterceptor(authRepo)),
//		)),
//	)
//
//	v1Example.RegisterExampleServiceServer(s, v1ExampleSrv.NewExampleServiceServer())
//
//	// SignUp Tag on the server.
//	v1tagOpts := &v1TagSrv.Options{
//		Debug: false,
//	}
//	// SignUp xx on the same server.
//	v1Tag.RegisterTagServiceServer(s, v1TagSrv.NewTagServiceServer(v1tagOpts, db.DB))
//
//
//	auth.RegisterAuthServiceServer(s, authSrv.NewAuthServiceServer(db.DB, jwt))
//	feedback.RegisterFeedbackServiceServer(s, feedbackSrv.NewFeedbackServiceServer(db.DB))
//	userTag.RegisterUserTagServiceServer(s, userTagSrv.NewUserTagServiceServer(db.DB))
//	userPost.RegisterUserPostServiceServer(s, userPostSrv.NewUserPostServiceServer(db.DB))
//
//	// run GRPC server
//	go func() {
//		if err = s.Serve(lis); err != nil {
//			log.Fatalf("failed to serve: %v", err)
//		}
//	}()
//
//	// run HTTP gateway
//	err = gateway.Run("localhost:50051", "127.0.0.1:8051")
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
//
//
