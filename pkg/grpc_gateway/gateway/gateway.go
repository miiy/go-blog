package gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

type RegisterServiceHandler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

// Run runs the gRPC-Gateway, dialling the provided address.
func Run(dialAddr, gatewayAddr string, handlers []RegisterServiceHandler, openAPIFS fs.FS) error {

	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	// https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_your_gateway/
	gwMux := runtime.NewServeMux(
		// runtime.WithForwardResponseOption(allowCORS),
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

	// err = pb.RegisterExampleServiceHandler(context.Background(), gwMux, conn)
	for _, v := range handlers {
		err = v(context.Background(), gwMux, conn)
		if err != nil {
			log.Fatalln("Failed to register gateway:", err)
		}
	}

	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/openapi") {
				r.URL.Path = strings.TrimPrefix(r.URL.Path, "/openapi")
				http.FileServer(http.FS(openAPIFS)).ServeHTTP(w, r)
				return
			}
			gwMux.ServeHTTP(w, r)
		}),
	}

	log.Printf("Serving gRPC-Gateway on http://%s", gatewayAddr)
	return gwServer.ListenAndServe()
}