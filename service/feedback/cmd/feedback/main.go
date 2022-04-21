package main

import (
	"flag"
	"fmt"
	feedbackpb "goblog.com/api/feedback/v1"
	"goblog.com/service/feedback/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func main()  {

	conf := flag.String("c", "./configs/default.yaml", "config file")
	port := flag.Int("addr", 50057, "the port to serve on")
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
	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(s, healthcheck)
	feedbackpb.RegisterFeedbackServiceServer(s, server.NewFeedbackServer(app.Database.DB))

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
