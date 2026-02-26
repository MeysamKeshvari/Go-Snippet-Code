package main

import "fmt"

func CheckType(t any) {
	switch t.(type) {
	case int:
		fmt.Println(t, "is integer")
	case string:
		fmt.Println(t, "is string")
	case float64:
		fmt.Println(t, "is float64")
	case bool:
		fmt.Println(t, "is bool")
	default:
		fmt.Println("Unknow type")
	}

}

func main() {
	b:=true
	CheckType(15)
	CheckType("Bob")
	CheckType(3.14)
	CheckType(b)
	CheckType(uint(5))
}