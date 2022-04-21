package main

import (
	"context"
	"flag"
	bookpb "goblog.com/api/book/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	var addr = flag.String("addr", "localhost:50053", "the address to connect to")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dit not connect: %v", err)
	}
	defer conn.Close()

	tc := bookpb.NewBookServiceClient(conn)

	log.Println("--- calling api.Book/Create ---")
	//callCreateArticle(tc)
	callGetBook(tc)
}

func callGetBook(client bookpb.BookServiceClient)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := bookpb.GetBookRequest{Id: 1}
	resp, err := client.GetBook(ctx, &req)
	if err != nil {
		log.Fatalf("client.GetBook(_) = _, %v", err)
	}
	log.Println("GetBook:", resp)
}