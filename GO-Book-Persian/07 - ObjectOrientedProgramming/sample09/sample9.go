package main

import "fmt"

type Stringr interface {
	String() string
}
type user struct {
	name string
}
type user2 struct {
	name string
}

func (u user) String() string {
	return "Hello, My Name is " + u.name
}
func PrintUserInfo(u any) {
	if v, ok := u.(Stringr); ok {
		fmt.Println(v.String())
	}else {
		fmt.Println("this object is not implemented Stringr interface")
	}
}

func main() {
	u:= user{"Bob"}
	u2:= user2{"Alice"}
	PrintUserInfo(u)
	PrintUserInfo(u2)
}