package main

import (
	"context"
	"flag"
	auth "github.com/miiy/go-blog/api/auth/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

type user struct {
	email string
	username string
	password string
}

func main() {
	var addr = flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("dit not connect: %v", err)
	}
	defer conn.Close()

	tc := auth.NewAuthServiceClient(conn)
	u1 := &user{
		email:    "a@a.com",
		username: "a3",
		password: "a3",
	}

	log.Println("--- calling up-api.Auth/SignUp ---")
	rResp, err := callSignUp(tc, u1.email, u1.username, u1.password, u1.password)
	if err != nil {
		log.Fatalf("client.SignUp(_) = _, %v", err)
	}
	log.Println("SignUp:", rResp)

	log.Println("--- calling up-api.Auth/SignIn ---")
	lResp, err := callSignIn(tc, u1.username, u1.password)
	if err != nil {
		log.Fatalf("client.SignIn(_) = _, %v", err)
	}
	log.Println("SignIn:", lResp)

	log.Println("--- calling up-api.Auth/VerifyToken ---")
	vResp, err := callVerifyToken(tc, lResp.AccessToken)
	if err != nil {
		log.Fatalf("client.VerifyToken(_) = _, %v", err)
	}
	log.Println("VerifyToken:", vResp)
}

func callSignUp(client auth.AuthServiceClient, email, username, password, passwordConfirmation string) (*auth.SignUpResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := auth.SignUpRequest{
		Email: email,
		Username: username,
		Password: password,
		PasswordConfirmation: passwordConfirmation,
	}
	return client.SignUp(ctx, &req)
}

func callVerifyToken(client auth.AuthServiceClient, accessToken string) (*auth.VerifyTokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := auth.VerifyTokenRequest{
		AccessToken: accessToken,
	}
	return client.VerifyToken(ctx, &req)
}

func callSignIn(client auth.AuthServiceClient, username, password string) (*auth.SignInResponse, error)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &auth.SignInRequest{
		Username: username,
		Password: password,
	}
	return client.SignIn(ctx, req)
}
