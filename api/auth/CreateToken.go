package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/waWanjohi/mux-auth/api/models"
)

// Generate your own secret key!
var secret = []byte("can-you-keep-a-secret?")

func CreateToken(sub string, userInfo interface{}) (string, error) {
	// Get the token instance with the Signing method
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// time frame for
	exp := time.Now().Add(time.Hour * 24)
	// Add your claims
	token.Claims = &models.MyJWTClaims{
		&jwt.RegisteredClaims{
			// Set the exp and sub claims. sub is usually the userID
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   sub,
		},
		userInfo,
	}
	// Sign the token with your secret key
	val, err := token.SignedString(secret)
	if err != nil {
		// On error return the error
		return "", err
	}
	// On success return the token string
	return val, nil
}
