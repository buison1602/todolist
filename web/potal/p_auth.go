package potal

import (
	"github.com/buison1602/todolist/storage"
	"golang.org/x/crypto/bcrypt"
)

type RegisterForm struct {
	Email        string `json:"email" validate:"required, email"`
	UserName     string `json:"username" validate:"required,alphanumeric,gte=4,lte=32"`
	PasswordHash string `json:"passwordHash" validate:"required"`
}

func (f *RegisterForm) FormCreate() (*storage.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(f.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &storage.User{
		Email:        f.Email,
		UserName:     f.UserName,
		PasswordHash: string(hash),
	}
	return user, nil
}

type LoginForm struct {
	UserName string `json:"username" validate:"required,alphanumeric,gte=4,lte=32"`
	Password string `json:"password" validate:"required"`
}
