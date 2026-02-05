package main

import (
	"fmt"
	"strconv"
	"errors"
)

func main(){
	//strToDigit()
	//digitToStr()

	err := checkPassword("test")
	if err!= nil {
		fmt.Println(err)
	}else {
		fmt.Println("welcome")
	}

	
}

func digitToStr(){
	str := strconv.Itoa(521)
	fmt.Println("your number is string now!", str)	
}

func strToDigit(){
	var(
		n int 
		err error 
		str = "Hello"
		//str = "360"
	)

	n ,err = strconv.Atoi(str)

	if err != nil {
		fmt.Println("your string does not contain number", err)
	} else {
		fmt.Print("Successfully converted!,", n)
	}	
}

func checkPassword(password string) error {
	if password != "somethings" {
		return errors.New("you are not authorized")
	}
	return nil 
}

