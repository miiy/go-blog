package main

import (
	"context"
	"flag"
	userpostpb "goblog.com/api/userpost/v1"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var addr = flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("dit not connect: %v", err)
	}
	defer conn.Close()

	tc := userpostpb.NewUserPostServiceClient(conn)

	log.Println("--- calling up-api.UserPost/Create ---")
	callCreate(tc)


}

func callCreate(client userpostpb.UserPostServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := userpostpb.CreateUserPostRequest{
		UserId:        1,
		Title:         "",
		Content:       "sss",
		Status:        0,
		PublishedTime: nil,
		UpdatedTime:   nil,
		Sort:          0,
	}
	resp, err := client.CreateUserPost(ctx, &req)
	if err != nil {
		log.Fatalf("client.Create(_) = _, %v", err)
	}
	log.Println("Create:", resp)
}
