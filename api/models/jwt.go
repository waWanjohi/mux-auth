package models

import "github.com/golang-jwt/jwt/v4"

type MyJWTClaims struct {
	*jwt.RegisteredClaims
	UserInfo interface{}
}

type ClaimsKey int
