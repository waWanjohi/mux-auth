package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func TokenValid(r *http.Request) error {
	// tokenString := ExtractToken(r)
	auth_header := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(auth_header, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)

	}
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	// b, err := json.MarshalIndent(data, "", " ")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(string(b))
}
