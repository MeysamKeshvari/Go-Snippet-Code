package main 

import "fmt"

func main(){
	f()
	fmt.Println("after panic")
}

func f(){
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("recoverd:", r)
		}
	}()
	panic("we are going to die!")
}
