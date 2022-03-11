package auth

import (
	"database/sql"
	"github.com/google/wire"
	pb "github.com/miiy/go-blog/api/auth/proto"
	authSrv "github.com/miiy/go-blog/api/auth/service"
	"github.com/miiy/go-blog/pkg/jwtauth"
)

func NewAuth(db *sql.DB, jwtAuth *jwtauth.JWTAuth) (pb.AuthServiceServer, error) {
	return authSrv.NewAuthServiceServer(db, jwtAuth), nil
}

var ProviderSet = wire.NewSet(NewAuth)