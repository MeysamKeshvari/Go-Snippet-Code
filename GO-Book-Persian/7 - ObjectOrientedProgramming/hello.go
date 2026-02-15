package main 

import "fmt"

func main(){

	u := User {
		FirstName : "Bob",
		LastName : "alice",
	}



	fmt.Println("Result" , u.GetFullName())
}

type User struct {
	FirstName string 
	LastName string
	fullName string
}

func (u *User) setFullName() {
	u.fullName = u.FirstName + " _" + u.LastName
}

func (u *User) GetFullName() string {
	u.setFullName()
	return u.fullName
}
  

