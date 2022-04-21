package main

import (
	"context"
	"flag"
	articlepb "goblog.com/api/article/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	var addr = flag.String("addr", "localhost:50052", "the address to connect to")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dit not connect: %v", err)
	}
	defer conn.Close()

	tc := articlepb.NewArticleServiceClient(conn)

	log.Println("--- calling api.Article/Create ---")
	//callCreateArticle(tc)
	callGetArticle(tc)

}

func callCreateArticle(client articlepb.ArticleServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := articlepb.CreateArticleRequest{
		Parent:    "parent/article",
		Article:   &articlepb.Article{
			Id:              0,
			UserId:          1,
			CategoryId:      1,
			Title:           "Title",
			MetaTitle:       "MetaTitle",
			MetaDescription: "MetaDescription",
			PublishedTime:   timestamppb.Now(),
			UpdatedTime:     timestamppb.Now(),
			FromText:        "FromText",
			FromUrl:         "FromUrl",
			Summary:         "Summary",
			Content:         "Content",
			Status:          articlepb.Article_ACTIVE,
		},
		ArticleId: "111111111111",
	}
	resp, err := client.CreateArticle(ctx, &req)
	if err != nil {
		log.Fatalf("client.CreateArticle(_) = _, %v", err)
	}
	log.Println("Create:", resp)
}

func callGetArticle(client articlepb.ArticleServiceClient)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := articlepb.GetArticleRequest{Id: 1}
	resp, err := client.GetArticle(ctx, &req)
	if err != nil {
		log.Fatalf("client.GetArticle(_) = _, %v", err)
	}
	log.Println("GetArticle:", resp)
}