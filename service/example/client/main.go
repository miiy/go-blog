package main

import (
	"context"
	"flag"
	v1Example "goblog.com/service/example/proto/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
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

	tc := v1Example.NewExampleServiceClient(conn)

	log.Println("--- calling up-api.Example/ExampleLog ---")
	callExampleLog(tc)
	log.Println("--- calling up-api.Example/ExamplePanic ---")
	callExamplePanic(tc)
	log.Println("--- calling up-api.Example/ReturnError1 ---")
	callExampleReturnError1(tc)
	log.Println("--- calling up-api.Example/ReturnError2 ---")
	callExampleReturnError2(tc)
	log.Println("--- calling up-api.Example/ReturnError3 ---")
	callExampleReturnError3(tc)
	log.Println("--- calling up-api.Example/callExampleValidate ---")
	callExampleValidate(tc)

}

func callExampleLog(client v1Example.ExampleServiceClient) *v1Example.ExampleResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Example.ExampleRequest{}
	resp, err := client.ExampleLog(ctx, &req)
	if err != nil {
		log.Fatalf("client.ExampleLog(_) = _, %v", err)
	}
	log.Println("ExampleLog:", resp)
	return resp
}

func callExamplePanic(client v1Example.ExampleServiceClient) *v1Example.ExampleResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Example.ExampleRequest{}
	resp, err := client.ExamplePanic(ctx, &req)
	if err != nil {
		log.Printf("client.ExamplePanic(_) = _, %v", err)
		return nil
	}
	log.Println("ExamplePanic:", resp)
	return resp
}

func callExampleReturnError1(client v1Example.ExampleServiceClient) *v1Example.ExampleResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Example.ExampleRequest{}
	resp, err := client.ExampleReturnError1(ctx, &req)
	if err != nil {
		log.Printf("client.ExampleReturnError1(_) = _, %v", err)
		return nil
	}
	log.Println("ExampleReturnError1:", resp)
	return resp
}

func callExampleReturnError2(client v1Example.ExampleServiceClient) *v1Example.ExampleResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Example.ExampleRequest{}
	resp, err := client.ExampleReturnError2(ctx, &req)
	if err != nil {
		log.Printf("client.ExampleReturnError2(_) = _, %v", err)
		return nil
	}
	log.Println("ExampleReturnError2:", resp)
	return resp
}

func callExampleReturnError3(client v1Example.ExampleServiceClient) *v1Example.ExampleResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Example.ExampleRequest{}
	resp, err := client.ExampleReturnError3(ctx, &req)
	if err != nil {
		s := status.Convert(err)
		log.Println(s)
		for _, d := range s.Details() {
			switch info := d.(type) {
			case *errdetails.BadRequest:
				log.Printf("Badrequest failure: %s", info)
			default:
				log.Printf("Unexpected type: %s", info)
			}
		}
		log.Printf("client.ExampleReturnError3(_) = _, %v", err)
		return nil
	}
	log.Println("ExampleReturnError3:", resp)
	return resp
}

func callExampleValidate(client v1Example.ExampleServiceClient) *v1Example.ExampleResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := v1Example.ExampleValidateRequest{
		Name: "",
		Age: 0,
	}
	resp, err := client.ExampleValidate(ctx, &req)
	if err != nil {
		log.Printf("client.ExampleReturnError2(_) = _, %v", err)
		return nil
	}
	log.Println("ExampleReturnError2:", resp)
	return resp
}