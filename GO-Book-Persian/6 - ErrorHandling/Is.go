package main 

import (
	"fmt"
	"errors"
)

func main(){

	err := RegisterUser()
	if err == invalidUserErr {
		fmt.Println("using == to find out if errors are equal")
	}

	if errors.is(err , invalidUserErr){
		fmt.Println("using errors.Is to find out if errors are equal")
	}
}

var invalidUserErr = errors.New("User is invalid")
func ValidateUser() error {
	return invalidUserErr
}

func RegisterUser() error{
	err:= ValidateUser()
	if err != nil {
		return fmt.Errorf("Error during registration: %w" , err)
	}
	return nil 
}
