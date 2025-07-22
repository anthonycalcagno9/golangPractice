package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	s2 := `password1234`
	bytes, err := bcrypt.GenerateFromPassword([]byte(s2), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hashed password:", string(bytes))

	err = bcrypt.CompareHashAndPassword(bytes, []byte(s2))
	if err != nil {
		fmt.Println("Password does not match:", err)
	} else {
		fmt.Println("Password matches!")
	}
}
