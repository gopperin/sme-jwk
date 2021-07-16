package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// TokenExpired 一些常量
var (
	ErrTokenExpired     error  = errors.New("Token is expired")
	ErrTokenNotValidYet error  = errors.New("Token not active yet")
	ErrTokenMalformed   error  = errors.New("That's not even a token")
	ErrTokenInvalid     error  = errors.New("Couldn't handle this token")
	ErrSignKey          string = "newtrekWang"
)

// User User
type User struct {
	KID      string `json:"kid"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Key      string `json:"key"`
}

// JWK JWK
type JWK struct {
	KID  string `json:"kid"`
	Keys string `json:"keys"`
}

// CustomClaims CustomClaims
type CustomClaims struct {
	UID   string `json:"uid"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Jti   string `json:"jti"`
	jwt.StandardClaims
}

// JWKKeys JWKKeys
type JWKKeys struct {
	Keys []JWKKey `json:"keys"`
}

// JWKKey JWKKey
type JWKKey struct {
	Alg string `json:"alg"`
	K   string `json:"k"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
}
