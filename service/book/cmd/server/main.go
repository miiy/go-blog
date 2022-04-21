package main

import (
	"flag"
	"fmt"
	bookpb "goblog.com/api/book/v1"
	"goblog.com/service/book/internal/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	conf := flag.String("c", "./configs/default.yaml", "config file")
	port := flag.Int("addr", 50053, "the port to serve on")
	flag.Parse()

	app, cleanUp, err := InitApplication(*conf)
	if err != nil {
		panic(err)
	}
	defer cleanUp()


	// gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Server listening at ", lis.Addr())

	s := grpc.NewServer()
	bookpb.RegisterBookServiceServer(s, server.NewBookServer(app.Database.Gorm, app.Logger))

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
