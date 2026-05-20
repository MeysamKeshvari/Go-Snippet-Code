package main

import "fmt"

func main(){
	fmt.Println("hello world")
	
	fmt.Println(`hello 
	world`)

	fmt.Println("Parsa is a good student")

	var age int 
	var name string
	var movie string 
	var score float64

	age = 13 
	name = "Parmis" 
	movie = "Stalag 17"
	score = 9.1

	var age_ = 20
	var _age int = 16
	var name_ = "Alice"
	var movie_ = "GOT"
	var score_ = 7.5

	var (
		jobTitle string = "Enginner"
		salary float64 = 2500.50
	)
	fmt.Println(jobTitle , salary)

	var (
		flowerName string
	)
	flowerName = "Rose"
	fmt.Println(flowerName)

	var firstName, lastName = "Bob","Boby"
	var firstName_ , lastName_ string = "Reza" , "Golzar"
	
	fmt.Println(firstName, lastName)
	fmt.Println(firstName_,lastName_)


	fmt.Println(name , "is a good student", "is", age , "years old", movie , "is favorite movie score is", score)
	fmt.Println(name_ , "is a good student", "is", age_ , "years old", movie_ , "is favorite movie score is", score_)
	fmt.Println(name_ , "is a good student", "is", _age , "years old", movie_ , "is favorite movie score is", score_)

	subject := "Policy"
	fmt.Println(subject)

	a,b,c := 1,2,3 
	fmt.Println(a,b)
	c = a 
	a = b 
	b = c 
	fmt.Println(a ,b)


	a1,a2 := 5,7
	a1,a2 = a2,a1
	fmt.Println(a1,a2)

	a3 := 3 
	b3 := 5
	maxCompare := max(a3 , b3)
	minCompare := min(a3,b3)
	fmt.Println("min", minCompare)
	fmt.Println("max" , maxCompare)


	var PI = 3.14 
	var castedPI int 
	fmt.Println("PI" ,PI)
	castedPI = int(PI)
	fmt.Println("casted: ",castedPI)

	fmt.Println(string(65),string(66),string(67))

	const pi = 3.14 

	const(
		 stateA = 1
		 stateB = 2 
		 stateC = 3
	)

	const(
		notPaid = iota + 1 
		paid 
		_
		canceled
	)


	fmt.Println("result const: ", canceled)

	str1 := "hello"
	str2 := "world"
	str3 := str1 +" "+ str2
	fmt.Println(str3)

	gender := "female"
	if gender == "female" {
		fmt.Println("You are woman!")
	} else {
		fmt.Println("You are man!")
	}

	n:=5 
	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	hand := "left"
	if hand == "left"{
		fmt.Println("hand left")
	} else if hand == "right"{
		fmt.Println("right")
	}
}