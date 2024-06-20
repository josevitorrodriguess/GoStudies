package main

import (
	"fmt"
	"module/packages"

	"github.com/badoux/checkmail"
)

func main() {
	packages.Typing()
	var email string

	fmt.Println("Type your email")
	fmt.Scan(&email)

	err  := checkmail.ValidateFormat(email)

	fmt.Println(err)
}

