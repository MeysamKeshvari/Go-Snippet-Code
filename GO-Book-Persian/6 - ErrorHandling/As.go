package main 


import(
	"errors"
	"fmt"
)

func main(){
	err := ValidateUser()
	validationErr := new(UserValidation)
	if errors.As(err ,validationErr){
		fmt.Println("Error:", validationErr.Message)
		fmt.Println("Email" , validationErr.Email)
		fmt.Println("code:" , validationErr.Code)
	} else {
		fmt.Println("It's not a UserValidation error type")
	}
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
		Code: 1234,
	}
}

