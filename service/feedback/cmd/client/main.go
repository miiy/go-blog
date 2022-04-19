package main

import (
	"flag"
	feedbackpb "goblog.com/api/feedback/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var addr = flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	// Set up a connection to the server.
	options := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(*addr, options...)
	if err != nil {
		log.Fatalf("dit not connect: %v", err)
	}
	defer conn.Close()

	_ = feedbackpb.NewFeedbackServiceClient(conn)

}
