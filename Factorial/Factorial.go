package main

import "fmt"

func main(){
	visualFactorial(3)
}

//3! = 3 * 2 * 1 = 6

func visualFactorial(n int) {
	
	for i := n; i <= n; i-- {
		n = n - 1
	}
	//middle := n-

	fmt.Println(n,"!", "=" ,n , "", "=" ,factorial(n))

}

func factorial(n int) int {
	if n == 0 {
		return 1 
	}
	return n * factorial(n-1)
}