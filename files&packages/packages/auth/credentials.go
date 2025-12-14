package auth

import "fmt"

// we can use function of this package in other packages if it is under package auth 


// start with small letter - unexported
// start with capital letter - exported

func LoginWithCreds(username string, password string) bool {
	fmt.Println("Logging in with username:", username)
	return true
}