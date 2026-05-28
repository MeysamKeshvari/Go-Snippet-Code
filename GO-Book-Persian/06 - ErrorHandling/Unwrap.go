package main 

import (
	"errors"
	"fmt"
)

var invalidUserErr = errors.New("User is invalid")

func ValidateUser() error {
	return invalidUserErr
}

func RegisterUser() error {
	err := ValidateUser()
	if err != nil {
		return fmt.Errorf("Error during registration: %w",err)
	}
	return nil
}

func main(){
	regErr := RegisterUser()
	fmt.Println(regErr)
	valErr := errors.Unwrap(regErr)
	fmt.Println(valErr)
}