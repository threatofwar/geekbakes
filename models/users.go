package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users = []User{
	{Username: "user1", Password: "password1"},
	{Username: "user2", Password: "password2"},
}
