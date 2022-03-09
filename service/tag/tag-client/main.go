package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	v1Tag "github.com/miiy/go-blog/service/tag/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

var commonAuthToken = ""

func ctxWithToken(ctx context.Context, scheme string, token string) context.Context {
	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", scheme, token))
	nCtx := metautils.NiceMD(md).ToOutgoing(ctx)
	return nCtx
}


func main() {
	var addr = flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(
		*addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("dit not connect: %v", err)
	}
	defer conn.Close()

	tc := v1Tag.NewTagServiceClient(conn)

	log.Println("--- calling up-api.Tag/Get ---")
	callTagList(tc, 0, 0)

	log.Println("--- calling up-api.Tag/Create ---")
	tag := callTagCreate(tc, "name", "description")

	log.Println("--- calling up-api.Tag/Get ---")
	callTagGet(tc, tag.Id)

	log.Println("--- calling up-api.Tag/Update ---")
	callTagUpdate(tc, tag.Id, "name1", "description1")

	log.Println("--- calling up-api.Tag/Get ---")
	callTagGet(tc, tag.Id)

	log.Println("--- calling up-api.Tag/Delete ---")
	callTagDelete(tc, tag.Id)

	log.Println("--- calling up-api.Tag/Get ---")
	callTagGet(tc, tag.Id)

	log.Println("--- calling up-api.Tag/List ---")
	callTagList(tc, 0, 0)
}

func callTagCreate(client v1Tag.TagServiceClient, name, description string) *v1Tag.TagId {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Tag.CreateTag{
		Name: name,
		Description: description,
	}
	resp, err := client.Create(ctx, &req)
	if err != nil {
		log.Fatalf("client.Create(_) = _, %v", err)
	}
	log.Println("Create:", resp)
	return resp
}

func callTagGet(client v1Tag.TagServiceClient, id int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Tag.TagId{
		Id: id,
	}
	resp, err := client.Get(ctx, &req)
	if err != nil {
		log.Fatalf("client.Get(_) = _, %v", err)
	}
	log.Println("Get:", resp)
}

func callTagUpdate(client v1Tag.TagServiceClient, id int64, name, description string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Tag.UpdateTag{
			Id: id,
			Name: name,
			Description: description,
	}
	resp, err := client.Update(ctx, &req)
	if err != nil {
		log.Fatalf("client.Update(_) = _, %v", err)
	}
	log.Println("Update:", resp)
}

func callTagDelete(client v1Tag.TagServiceClient, id int64)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Tag.TagId{
		Id: id,
	}
	resp, err := client.Delete(ctx, &req)
	if err != nil {
		log.Fatalf("client.Delete(_) = _, %v", err)
	}
	log.Println("Delete:", resp)
}

func callTagList(client v1Tag.TagServiceClient, page, perPage int64)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Tag.ListRequest{
		Page: page,
		PerPage: perPage,
	}
	resp, err := client.List(ctxWithToken(ctx, "bearer", commonAuthToken), &req)
	if err != nil {
		log.Fatalf("client.List(_) = _, %v", err)
	}
	log.Println("List:", resp)
}
