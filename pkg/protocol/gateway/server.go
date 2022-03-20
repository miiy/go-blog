//package gateway
//
//import (
//	"context"
//	"fmt"
//	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
//	auth "goblog.com/service/auth/proto"
//	example "goblog.com/service/example/proto/v1"
//	"goblog.com/service/tag/proto/v1"
//	userPost "goblog.com/service/userpost/proto"
//	userTag "goblog.com/service/usertag/proto"
//	"google.golang.org/grpc"
//	"google.golang.org/protobuf/encoding/protojson"
//	"google.golang.org/protobuf/proto"
//	"log"
//	"net/http"
//	"strings"
//)
//
//
//func Run(dialAddr, gatewayAddr string) error {
//	// Create a client connection to the gRPC Server we just started.
//	// This is where the gRPC-Gateway proxies the requests.
//	conn, err := grpc.DialContext(
//		context.Background(),
//		dialAddr,
//		grpc.WithBlock(),
//		grpc.WithInsecure(),
//	)
//	if err != nil {
//		return fmt.Errorf("failed to dial server: %w", err)
//	}
//
//
//	allowCORS := func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
//		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/")
//		w.Header().Set("Access-Control-Allow-Headers", "Origin,No-Cache, X-Requested-With, If-Modified-Since, Pragma, Last-Modified, Cache-Control, Expires, Content-Type, X-E4M-With, userId, token")
//		w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
//		return nil
//	}
//
//	// https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_your_gateway/
//	gwMux := runtime.NewServeMux(
//		runtime.WithForwardResponseOption(allowCORS),
//		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
//			MarshalOptions:   protojson.MarshalOptions{
//				UseProtoNames: true, // 使用proto中的大小写
//				EmitUnpopulated: true, // 显示未填充的字段
//			},
//			UnmarshalOptions: protojson.UnmarshalOptions{
//				DiscardUnknown: true,
//			},
//		}),
//	)
//
//	err = example.RegisterExampleServiceHandler(context.Background(), gwMux, conn)
//	if err != nil {
//		return fmt.Errorf("failed to register gateway: %w", err)
//	}
//	err = auth.RegisterAuthServiceHandler(context.Background(), gwMux, conn)
//	if err != nil {
//		return fmt.Errorf("failed to register gateway: %w", err)
//	}
//	err = v1.RegisterTagServiceHandler(context.Background(), gwMux, conn)
//	if err != nil {
//		return fmt.Errorf("failed to register gateway: %w", err)
//	}
//	err = userTag.RegisterUserTagServiceHandler(context.Background(), gwMux, conn)
//	if err != nil {
//		return fmt.Errorf("failed to register gateway: %w", err)
//	}
//	err = userPost.RegisterUserPostServiceHandler(context.Background(), gwMux, conn)
//	if err != nil {
//		return fmt.Errorf("failed to register gateway: %w", err)
//	}
//
//
//	srv := &http.Server{
//		Addr: gatewayAddr,
//		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			if strings.HasPrefix(r.URL.Path, "/api") {
//				gwMux.ServeHTTP(w, r)
//				return
//			}
//			// open api
//		}),
//	}
//
//	log.Printf("Serving gRPC-Gateway on http://%s\n", gatewayAddr)
//	return srv.ListenAndServe()
//}
//
//func t()  {
//
//	//log.Printf("Serving gRPC-Gateway on http://%s\n", gatewayAddr)
//	//return srv.ListenAndServe()
//
//	// graceful restart or stop
//
//	// Wait for interrupt signal to gracefully shutdown the server with
//	// a timeout of 5 seconds.
//	//quit := make(chan os.Signal)
//	//// kill (no param) default send syscanll.SIGTERM
//	//// kill -2 is syscall.SIGINT
//	//// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
//	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	//go func() {
//	//	<-quit
//	//	log.Println("Shutdown Server ...")
//	//
//	//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	//	defer cancel()
//	//	if err = srv.Shutdown(ctx); err != nil {
//	//		log.Fatal("Server Shutdown:", err)
//	//	}
//	//	// catching ctx.Done(). timeout of 5 seconds.
//	//	select {
//	//	case <-ctx.Done():
//	//		log.Println("timeout of 5 seconds.")
//	//	}
//	//	log.Println("Server exiting")
//	//}()
//
//	//
//	//// graceful shutdown
//	//c := make(chan os.Signal, 1)
//	//signal.Notify(c, os.Interrupt)
//	//go func() {
//	//	for range c {
//	//		// sig is a ^C, handle it
//	//	}
//	//
//	//	_, cancel := context.WithTimeout(ctx, 5*time.Second)
//	//	defer cancel()
//	//
//	//	if err := gwServer.Shutdown(ctx); err != nil {
//	//		log.Fatal("Server Shutdown:", err)
//	//	}
//	//}()
//	//
//	//log.Println("starting HTTP/REST gateway...")
//	//
//	//log.Printf("Serving gRPC-Gateway on http://%s\n", gatewayAddr)
//}
