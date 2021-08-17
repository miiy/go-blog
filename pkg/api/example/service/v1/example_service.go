package v1

import (
	"context"
	"errors"
	v1 "github.com/miiy/go-blog/pkg/api/example/proto/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"log"
)

type ExampleServiceServer struct {
	v1.UnimplementedExampleServiceServer
}

func NewExampleServiceServer() v1.ExampleServiceServer {
	return &ExampleServiceServer{
	}
}

func (s *ExampleServiceServer) ExampleLog(ctx context.Context, request *v1.ExampleRequest) (*v1.ExampleResponse, error) {
	grpclog.Infoln("info log")
	grpclog.Warningln("warning log")
	grpclog.Errorln("error log")
	return &v1.ExampleResponse{
	}, nil
}

func (s *ExampleServiceServer) ExamplePanic(ctx context.Context, request *v1.ExampleRequest) (*v1.ExampleResponse, error) {
	panic("panic error")
	return nil, nil
}

func (s *ExampleServiceServer) ExampleReturnError1(ctx context.Context, request *v1.ExampleRequest) (*v1.ExampleResponse, error) {
	st := status.New(codes.NotFound, "not found")
	return nil, st.Err()
}

func (s *ExampleServiceServer) ExampleReturnError2(ctx context.Context, request *v1.ExampleRequest) (*v1.ExampleResponse, error) {
	return nil, errors.New("error2")
}

func (s *ExampleServiceServer) ExampleReturnError3(ctx context.Context, request *v1.ExampleRequest) (*v1.ExampleResponse, error) {
	st := status.New(codes.InvalidArgument, "invalid argument")

	ds, err := st.WithDetails(&errdetails.BadRequest{
		FieldViolations: []*errdetails.BadRequest_FieldViolation{
			{
				Field:       "someRequest.email_address",
				Description: "invalid email address",
			},
		},
	})
	if err != nil {
		return nil, st.Err()
	}
	return nil, ds.Err()
}

func (s *ExampleServiceServer) ExampleValidate(ctx context.Context, request *v1.ExampleValidateRequest) (*v1.ExampleResponse, error) {
	log.Println(request)
	return &v1.ExampleResponse{}, nil
}