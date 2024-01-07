package helper

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	Id       int
	UserName string
}

func (a AuthClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthClaims) GetNotBefore() (*jwt.NumericDate, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthClaims) GetIssuer() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthClaims) GetSubject() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthClaims) GetAudience() (jwt.ClaimStrings, error) {
	//TODO implement me
	panic("implement me")
}
