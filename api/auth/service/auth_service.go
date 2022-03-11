package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	pb "github.com/miiy/go-blog/api/auth/proto"
	"github.com/miiy/go-blog/api/auth/repository"
	"github.com/miiy/go-blog/pkg/jwtauth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	repository *repository.Repository
	jwtAuth *jwtauth.JWTAuth
	log grpclog.LoggerV2
	pb.UnimplementedAuthServiceServer
}

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrPasswordsDiffer = errors.New("passwords differ")
	ErrUnauthenticated = errors.New("unauthenticated")

	ErrUsernameOrEmailExist = errors.New("username or email already exist")

	ErrUserNotFound  = errors.New("user not found")
	ErrWrongPassword = errors.New("wrong password")

)

// NewAuthServiceServer
//
func NewAuthServiceServer(db *sql.DB, jwtAuth *jwtauth.JWTAuth) pb.AuthServiceServer {
	r := repository.NewRepository(db)
	return &AuthServiceServer{
		repository: r,
		jwtAuth: jwtAuth,
	}
}

// AuthFuncOverride is called instead of defaultAuthFunc
//
func (s *AuthServiceServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	allowMethod := []string{
		"SignUp",
		"SignIn",
		"UsernameCheck",
		"EmailCheck",
		"PhoneCheck",
	}

	for _, method := range allowMethod {
		if fullMethodName == "/" + pb.AuthService_ServiceDesc.ServiceName + "/" + method {
			return ctx, nil
		}
	}

	return nil, status.New(codes.Unauthenticated, ErrUnauthenticated.Error()).Err()
}

func validateSignUp(request *pb.SignUpRequest) error {
	// validate
	if request.Username == "" || request.Email == "" || request.Password == "" || request.PasswordConfirmation == "" {
		return ErrInvalidArgument
	}
	if request.Password != request.PasswordConfirmation {
		return ErrPasswordsDiffer
	}
	return nil
}

// SignUp
//
func (s *AuthServiceServer) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	if err := validateSignUp(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	p := repository.RegisterParam{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}

	hashPasswd, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Errorln(err)
		return nil, err
	}
	p.Password = string(hashPasswd)

	exist, err := s.repository.UserExist(ctx, p)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, status.Error(codes.AlreadyExists, ErrUsernameOrEmailExist.Error())
	}

	// register
	id, err := s.repository.SignUp(ctx, p)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.FirstById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponse{
		User: &pb.SignUpResponse_User{
			Username: user.Username,
		},
	}, nil
}

// UsernameCheck
//
func (s *AuthServiceServer) fieldCheck(ctx context.Context, field, value string) (*pb.FieldCheckResponse, error) {
	exist, err := s.repository.FieldExist(ctx, field, value)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.FieldCheckResponse{
		Exist: exist,
	}, nil
}

// UsernameCheck
//
func (s *AuthServiceServer) UsernameCheck(ctx context.Context, request *pb.FieldCheckRequest) (*pb.FieldCheckResponse, error) {
	return s.fieldCheck(ctx, repository.FieldUsername, request.Value)
}

// EmailCheck
//
func (s *AuthServiceServer) EmailCheck(ctx context.Context, request *pb.FieldCheckRequest) (*pb.FieldCheckResponse, error) {
	return s.fieldCheck(ctx, repository.FieldEmail, request.Value)
}

// PhoneCheck
//
func (s *AuthServiceServer) PhoneCheck(ctx context.Context, request *pb.FieldCheckRequest) (*pb.FieldCheckResponse, error) {
	return s.fieldCheck(ctx, repository.FieldPhone, request.Value)
}

// validateSignIn
//
func validateSignIn(request *pb.SignInRequest) error {
	if request.Username == "" || request.Password == "" {
		return ErrInvalidArgument
	}
	return nil
}

// SignIn
//
func (s *AuthServiceServer) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	// Add fields the ctxtags of the request which will be added to all extracted loggers.
	grpc_ctxtags.Extract(ctx).Set("custom_tags.string", "something").Set("custom_tags.int", 1337)
	l := ctxzap.Extract(ctx)
	l.Info("zap log")

	if err := validateSignIn(request); err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	user, err := s.repository.FirstByUsername(ctx, request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		s.log.Error(err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, ErrWrongPassword
	}

	token, err := s.jwtAuth.CreateToken(user.Username)
	if err != nil {
		return nil, err
	}

	return &pb.SignInResponse{
		TokenType: "Bearer",
		AccessToken: token,
		ExpiresIn: s.jwtAuth.Options.ExpiresIn,
		User: &pb.SignInResponse_User{
			Username: user.Username,
		},
	}, nil
}

// validateVerifyToken
//
func validateVerifyToken(request *pb.VerifyTokenRequest) error {
	if request.AccessToken == "" {
		return ErrInvalidArgument
	}
	return nil
}

// VerifyToken
//
func (s *AuthServiceServer) VerifyToken(_ context.Context, request *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {

	if err := validateVerifyToken(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	claims, err := s.jwtAuth.ParseToken(request.AccessToken)
	if err != nil {
		if err == jwtauth.ErrTokenExpired {
			return nil, status.New(codes.NotFound, err.Error()).Err()
		}
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.VerifyTokenResponse{
		User: &pb.VerifyTokenResponse_User{
			Username: claims.Username,
		},
	}, nil
}

// RefreshToken
//
func (s *AuthServiceServer) RefreshToken(ctx context.Context, request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return &pb.RefreshTokenResponse{

	}, nil
}

// SignOut
//
func (s *AuthServiceServer) SignOut(ctx context.Context, request *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	return &pb.SignOutResponse{

	}, nil
}
