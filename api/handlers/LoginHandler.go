package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/waWanjohi/mux-auth/api/auth"
	"github.com/waWanjohi/mux-auth/api/models"
)

var LoginHandler = func(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("static/login.html"))

	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		template.Must(template.ParseFiles("static/index.html"))
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var u models.UserAuth
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}

	// Create user if your conditions match. Below, all username and passwords are accepted.
	user := &models.User{
		Id:       "2332-abcd-acdf-ccd2",
		Name:     "JWT Master",
		Username: string(u.Username),
		Password: string(u.Password),
	}

	tokenString, _ := auth.CreateToken(user.Id, user)
	payload := make(map[string]string)
	payload["access_token"] = tokenString

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(payload)

}
