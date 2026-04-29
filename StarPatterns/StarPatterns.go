package main 

import (
	"fmt"
	"strings"
)

func main(){
	star()
}

func starPattern(){
	for i:=1; i<=5; i++ {
		for j:=1; j<=i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

func starPatternRight(){
	for i:=5; i>=1; i-- {
		for j:=1; j<=i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
		
	}
}

func starPatterns(){
	for i:=1; i<= 5; i++ {
		for j:=1; j<=i; j++ {
			fmt.Print("*\t" )
		}
		fmt.Println()
	}
}

func star(){
	for i:=1; i<=50; i++ {
		for j:=1; j<=i; j++ {			
			if i==j {
				fmt.Print(strings.Repeat(" ",i),"*")
			}
		}
		fmt.Println()
    }
}

