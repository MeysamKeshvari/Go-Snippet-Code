package main

import (
	"fmt"
	"strconv"
)

func main() {
	strToDigit()
}

func strToDigit() {
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
