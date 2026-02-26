package main 

import "fmt"

func main(){
	f()
	fmt.Println("after panic")
}

func f(){
	panic("we are going to die!")
}
