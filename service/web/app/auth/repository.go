package auth

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) firstByUsername(username string) *User {
	user := &User{} // var user = new(User)
	if err := r.db.Where("username=?", username).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		log.Fatal(err)
		return nil
	}
	return user
}

func (r *repository) firstByUsernamePassword(username, password string) *User {
	user := &User{}
	if err := r.db.Where("username=? and password=?", username, password).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		log.Fatal(err)
		return nil
	}
	return user
}

func (r *repository) create(user *User) (*User, error) {
	err := r.db.Create(user).Error
	return user, err
}

func (r *repository) updates(user *User) (*User, error) {
	err := r.db.Updates(user).Error
	return user, err
}

func (r *repository) delete(user *User) (*User, error) {
	err := r.db.Delete(user).Error
	return user, err
}