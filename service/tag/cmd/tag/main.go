package main

import (
	"flag"
	"fmt"
	tagpb "goblog.com/api/tag/v1"
	"goblog.com/service/tag/internal/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {

	conf := flag.String("c", "./configs/default.yaml", "config file")
	port := flag.Int("addr", 50054, "the port to serve on")
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
	o := &server.Options{Debug: true}
	tagpb.RegisterTagServiceServer(s, server.NewTagServiceServer(o, app.Database.DB))

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
