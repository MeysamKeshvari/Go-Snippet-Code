package main

import (
	"fmt"
	"errors"
)

func main(){
	
	err := checkPassword("test")
	if err!= nil {
		fmt.Println(err)
	}else {
		fmt.Println("welcome")
	}
}

func checkPassword(password string) error {
	if password != "somethings" {
		return errors.New("you are not authorized")
	}
	return nil 
}