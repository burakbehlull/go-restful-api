package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	Age      int    `json:"age"`
}
