package auth

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type key int

const (
	keyPrincipalID key = iota
)

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func SetJWTClaimsContext(ctx context.Context, claims jwt.MapClaims) context.Context {
	return context.WithValue(ctx, keyPrincipalID, claims)
}

func JWTClaimsFromContext(ctx context.Context) (jwt.MapClaims, bool) {
	claims, ok := ctx.Value(keyPrincipalID).(jwt.MapClaims)
	return claims, ok
}
