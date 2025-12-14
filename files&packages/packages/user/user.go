package user

// import "fmt"

type User struct {
	Username string
	Password string
}

var users []User

func NewUser(us string, pas string) *User {
	users = append(users, User{Username: us, Password: pas})
	return &User{
		Username: us,
		Password: pas,
	}
}

func AllUsers()  []User {
	// fmt.Println("Fetching all users", len(users))
	// fmt.Println(users)
	return users
}