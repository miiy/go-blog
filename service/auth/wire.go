package auth

import (
	"database/sql"
	"github.com/google/wire"
	pb "goblog.com/service/auth/proto"
	authSrv "goblog.com/service/auth/service"
	"goblog.com/pkg/jwtauth"
)

func NewAuth(db *sql.DB, jwtAuth *jwtauth.JWTAuth) (pb.AuthServiceServer, error) {
	return authSrv.NewAuthServiceServer(db, jwtAuth), nil
}

var ProviderSet = wire.NewSet(NewAuth)