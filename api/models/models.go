package models

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MyJWTClaims struct {
	*jwt.RegisteredClaims
	UserInfo interface{}
}

type ClaimsKey int
