package main

import "fmt"

func main(){
	showDetail(pizza)
	
}

func pizza()(string , int ,string){
	return "PiZZA" , 3500 , "Cheese + PePeroni"
}

func showDetail(f func() (string,int,string)){
	name,price,ingredients := f()
	fmt.Println("Show food details:")
	fmt.Println("Name:" , name)
	fmt.Println("Price:" , price)
	fmt.Println("ingredients:" , ingredients)
}