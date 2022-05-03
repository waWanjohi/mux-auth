package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/waWanjohi/mux-auth/api/auth"
)

var ProtectedHandler = func(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.JWTClaimsFromContext(r.Context())

	//Do something with the UserInfo claims
	if val, ok := claims["UserInfo"]; ok {
		userinfo := val.(map[string]string)
		fmt.Print(userinfo)
	}
	//Do something with the sub claim
	if val, ok := claims["sub"]; ok {
		userid := val.(string)
		fmt.Print(userid)
	}

	if !ok {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(claims)
}
