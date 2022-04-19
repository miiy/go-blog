package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type service struct {
	repository *repository
}

var (
	errUserNotFound = errors.New("username or password is error")
	errInternalServerError = errors.New("internal server error")
)

func NewService(repository *repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) signUp(p *SignUpParam) (*APIUser, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.create(&User{
		Email: p.Email,
		Username: p.Username,
		Password: string(hash),
	})
	if err != nil {
		return nil, err
	}

	return &APIUser{
		Username: user.Username,
	}, nil
}

func (s *service) signIn(p *SignInParam) (*User, error){
	user := s.repository.firstByUsername(p.Username)
	if user == nil {
		return nil, errUserNotFound
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	return user, err
}

func createJwtToken(username string) (string, error) {
		token, err := module.jwtAuth.CreateToken(username)
		if err != nil {
			log.Println(err)
			return "", errInternalServerError
		}
		return token, nil
}
