package auth

import (
	"database/sql"
	"github.com/google/wire"
	authpb "goblog.com/api/auth/v1"
	"goblog.com/pkg/jwtauth"
	authSrv "goblog.com/service/auth/service"
)

func NewAuth(db *sql.DB, jwtAuth *jwtauth.JWTAuth) (authpb.AuthServiceServer, error) {
	return authSrv.NewAuthServiceServer(db, jwtAuth), nil
}

var ProviderSet = wire.NewSet(NewAuth)