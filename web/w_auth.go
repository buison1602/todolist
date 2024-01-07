package web

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/asaskevich/govalidator"
	"github.com/buison1602/todolist/helper"
	"github.com/buison1602/todolist/web/potal"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	var f potal.RegisterForm
	err := parseJSON(r, &f)
	if err != nil {
		response(w, http.StatusBadRequest, nil, err)
		return
	}

	if govalidator.IsNull(f.Email) || govalidator.IsNull(f.PasswordHash) || govalidator.IsNull(f.UserName) {
		response(w, http.StatusBadRequest, nil, helper.DataError)
		return
	}
	if !govalidator.IsEmail(f.Email) {
		response(w, http.StatusBadRequest, nil, helper.EmailError)
		return
	}

	user, err := f.FormCreate()
	if err != nil {
		response(w, http.StatusInternalServerError, nil, err)
		return
	}
	err = s.db.CheckDuplicate(user)
	if err != nil {
		response(w, http.StatusInternalServerError, nil, err)
		return
	}

	err = s.db.Create(&user)
	if err != nil {
		response(w, http.StatusInternalServerError, nil, err)
		return
	}
	response(w, http.StatusCreated, Map{
		"userId": user.Id,
	}, nil)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var f potal.LoginForm
	err := parseJSON(r, &f)
	if err != nil {
		response(w, http.StatusBadRequest, nil, err)
		return
	}
	if govalidator.IsNull(f.Password) || govalidator.IsNull(f.UserName) {
		response(w, http.StatusBadRequest, nil, helper.DataError)
		return
	}

	user, err := s.db.FirstByUserName(f.UserName)
	if err != nil {
		response(w, http.StatusBadRequest, nil, err)
		return
	}
	// compare sent in pass with saved user pass hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(f.Password))
	if err != nil {
		response(w, http.StatusBadRequest, nil, helper.LoginError)
		return
	}

	// Create a private key of type ECDSA
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		response(w, http.StatusInternalServerError, nil, err)
		return
	}

	// generate a jwt token
	var authClaim = helper.AuthClaims{
		Id:       user.Id,
		UserName: user.UserName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, authClaim)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		response(w, http.StatusInternalServerError, nil, err)
		return
	}

	response(w, http.StatusOK, Map{
		"token": tokenString,
		"user":  user,
	}, nil)
}
