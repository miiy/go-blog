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

type Options struct {
	// Addr is the address to listen
	Addr string
	// GRPCServer defines an endpoint of a gRPC service
	GRPCServerAddr string
	// OpenAPIFS
	OpenAPIFS fs.FS
	// ServerMuxOption
	ServerMuxOption []runtime.ServeMuxOption
	// RegisterServiceHandlers
	RegisterServiceHandlers []RegisterServiceHandler
}

// Run runs the gRPC-Gateway, dialling the provided address.
func Run(ctx context.Context, opts Options) error {

	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		ctx,
		opts.GRPCServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	// https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_your_gateway/
	serverMuxOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions:   protojson.MarshalOptions{
			UseProtoNames: true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	opts.ServerMuxOption = append(opts.ServerMuxOption, serverMuxOption)
	gwMux := runtime.NewServeMux(opts.ServerMuxOption...)

	// err = pb.RegisterExampleServiceHandler(context.Background(), gwMux, conn)
	for _, f := range opts.RegisterServiceHandlers {
		if err = f(ctx, gwMux, conn); err != nil {
			log.Fatalln("Failed to register gateway:", err)
		}
	}

	gwServer := &http.Server{
		Addr: opts.Addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/openapi") {
				r.URL.Path = strings.TrimPrefix(r.URL.Path, "/openapi")
				http.FileServer(http.FS(opts.OpenAPIFS)).ServeHTTP(w, r)
				return
			}
			gwMux.ServeHTTP(w, r)
		}),
	}

	log.Printf("Serving gRPC-Gateway on http://%s", opts.Addr)
	return gwServer.ListenAndServe()
}