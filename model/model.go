package model

import "database/sql"

type User struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

var users = []User{
	{
		Username: "user1",
		FullName: "User One",
		Email:    "user1@example.com",
		IsActive: true,
	},
	{
		Username: "user2",
		FullName: "User Two",
		Email:    "user2@example.com",
	},
}

var countries = []Country{
	{Name: "Nigeria"},
	{Name: "Ghana"},
	{Name: "Gambia"},
	{Name: "South Africa"},
	{Name: "Argentina"},
	{Name: "Chile"},
	{Name: "Peru"},
	{Name: "Colombia"},
	{Name: "Venezuela"},
}

type Country struct {
	Name string `json:"name"`
}

func ListUsers() []User {
	return users
}

func GetUser(username string) (*User, error) {
	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, sql.ErrNoRows
}

func CreateUser(user User) error {
	users = append(users, user)
	return nil
}

func ListCountries() []Country {
	return countries
}
