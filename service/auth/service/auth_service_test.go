package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	authpb "goblog.com/api/auth/v1"
	"goblog.com/pkg/jwtauth"
	"log"
	"testing"
)

var (
	mock sqlmock.Sqlmock
)

func newSrv() (authpb.AuthServiceServer, func()) {
	var db *sql.DB
	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	jwt := jwtauth.NewJWTAuth(&jwtauth.Options{
		Secret:    "abcd1234",
		ExpiresIn: 100,
	})

	srv := NewAuthServiceServer(db, jwt)
	closeFunc := func() {
		db.Close()
	}
	return srv, closeFunc
}

func TestAuthServiceServer_validateRegister(t *testing.T) {
	jsonData := `
{
	"success": [
		{"email": "test@email.com", "username": "test", "password": "test", "password_confirmation": "test"}
	],
	"fail": [
		{"email": "",               "username": "test", "password": "test", "password_confirmation": "test"},
		{"email": "test@email.com", "username": "",     "password": "test", "password_confirmation": "test"},
		{"email": "test@email.com", "username": "test", "password": "",     "password_confirmation": "test"},
		{"email": "test@email.com", "username": "test", "password": "test", "password_confirmation": ""},
		{"email": "test@email.com", "username": "test", "password": "test", "password_confirmation": "test123"}
	]
}
`
	var data map[string][]*authpb.SignUpRequest
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data["success"] {
		err := validateSignUp(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, v := range data["fail"] {
		err := validateSignUp(v)
		if err == nil {
			log.Fatal(err)
		}
	}
}

func TestAuthServiceServer_Register(t *testing.T) {

	srv, closeFunc := newSrv()
	defer closeFunc()

	ctx := context.Background()

	t.Run("test validate", func(t *testing.T) {
		_, err := srv.SignUp(ctx, &authpb.SignUpRequest{
			Email:                "",
			Username:             "t",
			Password:             "t",
			PasswordConfirmation: "t",
		})
		if err == nil {
			t.Fatal(err)
		}
		t.Log(err)
	})

	mock.ExpectPrepare("SELECT (.+) FROM users")
	cRows := mock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT (.+) FROM users").WithArgs("t", "t").WillReturnRows(cRows)

	mock.ExpectPrepare("INSERT INTO users")
	mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 0))

	mock.ExpectPrepare("SELECT (.+) FROM users")
	uRows := mock.NewRows([]string{"id", "username", "password", "status"}).AddRow(1, "t", "", 0)
	mock.ExpectQuery("SELECT (.+) FROM users").WithArgs(1).WillReturnRows(uRows)

	res, err := srv.SignUp(ctx, &authpb.SignUpRequest{
		Email:                "t",
		Username:             "t",
		Password:             "t",
		PasswordConfirmation: "t",
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Println(res)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAuthServiceServer_validateSignIn(t *testing.T) {
	jsonData := `
{
	"success": [
		{"username": "test", "password": "test"}
	],
	"fail": [
		{"username": "", "password": ""}
		{"username": "test", "password": ""}
		{"username": "test", "password": "test123"}
	]
}
`
	var data map[string][]*authpb.SignInRequest
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatal(err)
	}
}