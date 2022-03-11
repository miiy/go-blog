package auth

import (
	"encoding/gob"
	"github.com/miiy/go-web/pkg/database"
	"gorm.io/gorm"
)

//type CommonColumn database.CommonColumn

// Model
type User struct {
	gorm.Model
	Username        string            `json:"username" gorm:"uniqueIndex"`
	Password        string            `json:"password"`
	Email           string            `json:"email" gorm:"uniqueIndex"`
	EmailVerifiedAt database.JSONTime `json:"email_verified_at"`
	Phone           string            `json:"phone" gorm:"uniqueIndex"`
	Status int8                       `json:"status"`
}

// Session
type AuthUser struct {
	Username string
}

// API
type APIUser struct {
	Username string `json:"username"`
}

// signIn
type APIUserSigIn struct {
	TokenType   string   `json:"token_type"`
	AccessToken string   `json:"access_token"`
	ExpiresIn   int64    `json:"expires_in"`
	User        *APIUser `json:"user"`
}

// Validation

type SignUpParam struct {
	Email                string `json:"email" form:"email" binding:"required,email"`
	Username             string `json:"username" form:"username" binding:"required,gte=4,lte=30,is_exists"`
	Password             string `json:"password" form:"password" binding:"required,gte=6"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" binding:"required,eqfield=Password"`
}

type SignInParam struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RefreshTokenParam struct {
	AccessToken string `json:"access_token" binding:"required"`
}


func init()  {
	gob.Register(&AuthUser{})
}