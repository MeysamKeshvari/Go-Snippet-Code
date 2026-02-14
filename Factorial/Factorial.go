package main

import (
	"fmt"
	"strings"
)

func main(){
	factorialWithSteps(3)
}

func factorialWithSteps(n int){
	result := 1
	var parts []string
	for i := n; i >= 1; i-- {
		result *= i
		parts = append(parts, fmt.Sprintf("%d",i))
	}
	fmt.Printf("%d! = %s = %d\n", n, strings.Join(parts, " × "), result)
}

func factorial(n int) int {
	if n == 0 {
		return 1 
	}
	return n * factorial(n-1)
}


