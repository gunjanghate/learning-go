package main

import (
	"github.com/gunjanghate/learning-go/auth"
	"github.com/gunjanghate/learning-go/user"
	"github.com/fatih/color"
)

// DRY : Don't Repeat Yourself
// packages help to organize code into reusable modules

// we can use functions from other packages by importing them

func main() {
	auth.LoginWithCreds("gg", "gg1234")

	session := auth.GetSession()
	println("Session ID:", session)

	u := user.NewUser("gunjan", "gg1234")

	println("Created user:", u.Username)

	uss := user.AllUsers()

	println("Total users:", len(uss))
	// list all users

	for _, usr := range uss {
		// println("User:", usr.Username, usr.Password)
		color.Green("User: %s, Password: %s", usr.Username, usr.Password)
	}

}
