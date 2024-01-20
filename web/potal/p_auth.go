package potal

import (
	"github.com/buison1602/todolist/storage"
	"golang.org/x/crypto/bcrypt"
)

type RegisterForm struct {
	Email    string `json:"email" validate:"omitempty,email"`
	UserName string `json:"userName" validate:"required,alphanum,gte=4,lte=32"`
	Password string `json:"password" validate:"required"`
}

func (f *RegisterForm) FormCreate() (*storage.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)
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
	UserName string `json:"userName" validate:"required,alphanum,gte=4,lte=32"`
	Password string `json:"password" validate:"required"`
}
