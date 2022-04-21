package main

import (
	"context"
	"flag"
	bookpb "goblog.com/api/book/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	callListBooks(ctx, tc)
}


func callCreateBook(ctx context.Context,client bookpb.BookServiceClient) {
	req := bookpb.CreateBookRequest{
		Parent: "",
		Book:   &bookpb.Book{
			Id:              0,
			UserId:          0,
			CategoryId:      0,
			Name:            "test",
			Publisher:       "test",
			Year:            0,
			Pages:           0,
			Price:           0,
			Binding:         "test",
			Series:          "",
			Isbn:            "",
			BookDescription: "",
			AboutTheAuthor:  "",
			TableOfContents: "",
			Content:         "",
			Status:          0,
			CreateAt:      nil,
			UpdateAt:      nil,
			DeletedAt:     nil,
		},
		BookId: "",
	}
	resp, err := client.CreateBook(ctx, &req)
	if err != nil {
		log.Fatalf("client.CreateBook(_) = _, %v", err)
	}
	log.Println("CreateBook:", resp)
}

func callGetBook(ctx context.Context,client bookpb.BookServiceClient)  {
	req := bookpb.GetBookRequest{Id: 1}
	resp, err := client.GetBook(ctx, &req)
	if err != nil {
		s := status.Convert(err)
		if s.Code() == codes.NotFound {
			log.Println("not found")
		}
		log.Fatalf("client.GetBook(_) = _, %v", err)
	}
	log.Println("GetBook:", resp)
}

func callListBooks(ctx context.Context, client bookpb.BookServiceClient)  {
	req := bookpb.ListBooksRequest{
		Page:     1,
		PageSize: 10,
	}
	resp, err := client.ListBooks(ctx, &req)
	if err != nil {
		log.Fatalf("client.ListBooks(_) = _, %v", err)
	}
	log.Println("ListBooks:", resp)
}