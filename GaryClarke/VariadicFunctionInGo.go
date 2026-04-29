package main	

import "fmt"


func main(){
	sliceShareData()
	fmt.Println(sum(1,2,3))
	fmt.Println(sum(1,2,3,1))

	greet("Hello","Alice","Bob","john")

	fmt.Println(multiplyAll(5,2))
	fmt.Println(multiplyAll(5,3))
	fmt.Println(multiplyAll())
}

func sliceShareData(){
	a := []int {1,2,3}
	b := a 
	b[1] = 99
	fmt.Println(a)
}

func sum(numbers ...int)int {
	total := 0
	for _,n := range numbers {
		total += n 
	} 
	return total
}

func greet(prefix string,names ...string) {

	for _,name := range names {
		fmt.Println(prefix,name)
	}
}

func multiplyAll(nums ... int)int {

	if len(nums) == 0 {
		return 1 
	}

	result := 1 
	for _,n := range nums{
		result *= n
	}
	
	return result
}
