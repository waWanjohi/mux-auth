package models

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
