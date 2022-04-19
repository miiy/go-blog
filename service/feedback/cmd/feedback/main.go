package main

import (
	"flag"
	"fmt"
	feedbackpb "goblog.com/api/feedback/v1"
	"goblog.com/service/feedback/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {

	conf := flag.String("c", "./configs/default.yaml", "config file")
	addr := flag.String("addr", "50052", "0.0.0.0：50051")
	flag.Parse()

	app, cleanUp, err := InitApplication(*conf)
	if err != nil {
		panic(err)
	}
	defer cleanUp()


	// gRPC server
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	fmt.Println("Server listening at ", lis.Addr())

	s := grpc.NewServer()
	feedbackpb.RegisterFeedbackServiceServer(s, service.NewFeedbackServiceServer(app.Database.DB))
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
