package main

import (
	"context"
	"flag"
	userPost "github.com/miiy/go-blog/pkg/api/userpost/proto"
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

	tc := userPost.NewUserPostServiceClient(conn)

	log.Println("--- calling up-api.UserPost/Create ---")
	callCreate(tc)


}

func callCreate(client userPost.UserPostServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := userPost.CreateUserPost{
		UserId:        1,
		Title:         "",
		Content:       "sss",
		Status:        0,
		PublishedTime: nil,
		UpdatedTime:   nil,
		Sort:          0,
	}
	resp, err := client.Create(ctx, &req)
	if err != nil {
		log.Fatalf("client.Create(_) = _, %v", err)
	}
	log.Println("Create:", resp)
}
