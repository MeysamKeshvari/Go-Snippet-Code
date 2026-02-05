package main 

import "fmt"

func main(){
	err := ValidateUser()
	fmt.Println(err.Error())
}

type UserValidation struct {
	Code int 
	Email string
	Message string 
}

func (r UserValidation) Error() string {
	return r.Message
}

func ValidateUser() error {
	return UserValidation {
		Email: "Alice@gmail.com", 
		Message: "Email Already Exists",
		Code: 123,
	}
}

