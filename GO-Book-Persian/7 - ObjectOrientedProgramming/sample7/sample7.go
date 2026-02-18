package main

import "fmt"

func main() {
	Add(3,2)
}

func Add(a, b any) {
	inta:=a.(int)
	intb:=b.(int)
	fmt.Println(inta+intb)
}