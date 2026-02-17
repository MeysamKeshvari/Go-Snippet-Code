package main 

//import "fmt"

func main() {
	NewUser("alice","bob")
}

type user struct {
	firstName string
	lastName string 
}

func NewUser(first, last string) *user {
	return &user{
		firstName: first,
		lastName: last,
	}
}