package grpc

import (
	"context"
	"github.com/miiy/go-blog/service/tag/proto/v1"
	"google.golang.org/grpc"
	"net"
)

func RunServer(ctx context.Context, srv v1.TagServiceServer, address string) error {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterTagServiceServer(server, srv)

	return server.Serve(listen)
}
