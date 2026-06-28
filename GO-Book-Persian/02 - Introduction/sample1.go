package main

import (
	"fmt"
	"net"
	"log"
)


func main() {
	
	const (
		network = "tcp"
		address = "127.0.0.1:"
	)

	listener,err:= net.Listen(network,address)
	if err != nil {
		log.Fatalln("can't listen on given address",address,err)
	}
	fmt.Println("Listener address", listener.Addr())



	fmt.Println("hello world")

	fmt.Println(`hello 
	world`)

	fmt.Println("Alice is a good student")

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
		jobTitle string  = "Enginner"
		salary   float64 = 2500.50
	)
	fmt.Println(jobTitle, salary)

	var (
		flowerName string
	)
	flowerName = "Rose"
	fmt.Println(flowerName)

	var firstName, lastName = "Bob", "Boby"
	var firstName_, lastName_ string = "Reza", "Golzar"

	fmt.Println(firstName, lastName)
	fmt.Println(firstName_, lastName_)

	fmt.Println(name, "is a good student", "is", age, "years old", movie, "is favorite movie score is", score)
	fmt.Println(name_, "is a good student", "is", age_, "years old", movie_, "is favorite movie score is", score_)
	fmt.Println(name_, "is a good student", "is", _age, "years old", movie_, "is favorite movie score is", score_)

	subject := "Policy"
	fmt.Println(subject)

	a, b, c := 1, 2, 3
	fmt.Println(a, b)
	c = a
	a = b
	b = c
	fmt.Println(a, b)

	a1, a2 := 5, 7
	a1, a2 = a2, a1
	fmt.Println(a1, a2)

	a3 := 3
	b3 := 5
	maxCompare := max(a3, b3)
	minCompare := min(a3, b3)
	fmt.Println("min", minCompare)
	fmt.Println("max", maxCompare)

	var PI = 3.14
	var castedPI int
	fmt.Println("PI", PI)
	castedPI = int(PI)
	fmt.Println("casted: ", castedPI)

	fmt.Println(string(65), string(66), string(67))

	const pi = 3.14

	const (
		stateA = 1
		stateB = 2
		stateC = 3
	)

	const (
		notPaid = iota + 1
		paid
		_
		canceled
	)

	fmt.Println("result const: ", canceled)

	str1 := "hello"
	str2 := "world"
	str3 := str1 + " " + str2
	fmt.Println(str3)

	gender := "female"
	if gender == "female" {
		fmt.Println("You are woman!")
	} else {
		fmt.Println("You are man!")
	}

	n := 5
	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	hand := "left"
	if hand == "left" {
		fmt.Println("hand left")
	} else if hand == "right" {
		fmt.Println("right")
	}

	day := 4
	switch day {
	case 0:
		fmt.Println("Sat")
	case 1:
		fmt.Println("Sun")
	case 3:
		fmt.Println("Mon")
	case 4:
		fmt.Println("Tue")
	default:
		fmt.Println("Invalid")
	}

	n1 := 5
	switch n1 {
	case 4:
		fmt.Println(4)
		fallthrough
	case 5:
		fmt.Println(5)
		fallthrough
	case 6:
		fmt.Println(6)
		fallthrough
	case 7:
		fmt.Println(7)
	case 8:
		fmt.Println(8)
	case 9:
		fmt.Println(9)
	}

	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}

	fmt.Println("")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	i := 1
	for i < 10 {
		i++
		fmt.Println(i)
	}
	for i := 1; i <= 5; i++ {
		fmt.Print("A ")
		continue
		fmt.Print("B ")
	}

	fmt.Println()
	for i := 1; i <= 5; i++ {
		fmt.Println("A ")
		break
	}
	fmt.Println("loop finished")

	var i1 int
	for {
		i1++
		if i1 > 5 {
			break
		}
		fmt.Print(i1, " ")
	}

	var i2 int
	for i2 < 5 {
		i2++
		fmt.Print(i2, " ")
	}

	goto b
	fmt.Println("A")
b:
	fmt.Println("B")
}


