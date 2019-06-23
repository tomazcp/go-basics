package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := "password"
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	err = bcrypt.CompareHashAndPassword(bs, []byte(pw))
	if err != nil {
		fmt.Println("CAN'T LOGIN")
		return
	}
	fmt.Println("You're logged in")
}
