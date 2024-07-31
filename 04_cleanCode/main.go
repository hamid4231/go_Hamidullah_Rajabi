package main

import (
	"fmt"
)

// User struct represents a user with an email and password
type User struct {
	Email    string
	Password string
}

// UserRepo struct represents a repository of users
type UserRepository struct {
	database []User
}

// Register adds a new user to the repository
func (r *UserRepository) Register(u User) {
	if u.Email == "" || u.Password == "" {
		fmt.Println("Register failed: Email or Password cannot be empty")
		return
	}

	r.database = append(r.database, u)
	fmt.Println("User registered successfully")
}

// Login checks user credentials and returns an auth token if successful
func (r *UserRepository) Login(u User) string {
	if u.Email == "" || u.Password == "" {
		fmt.Println("Login failed: Email or Password cannot be empty")
		return ""
	}

	for _, user := range r.database {
		if user.Email == u.Email && user.Password == u.Password {
			return "auth token"
		}
	}

	fmt.Println("Login failed: Invalid email or password")
	return ""
}
