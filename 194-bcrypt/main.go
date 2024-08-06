package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("Original password:", s)
	fmt.Println("Hashed password:", string(bs))

	loginPass := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err != nil {
		fmt.Println("You can't login")
		return
	}
	fmt.Println("You're logged in")
}
