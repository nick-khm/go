package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level int
}

func main() {
	u := User{
		Name:  "Bob",
		Email: "bob@golang.com",
	}
	fmt.Printf("User=%v\n", u)

	a := Admin{Level: 2, User: User{Name: "Alice", Email: "alice@golang.com"}}
	fmt.Printf("Admin=%v\n", a)
}
