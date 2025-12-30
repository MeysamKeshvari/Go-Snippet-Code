package main  

import (
	"strings"
	"fmt"
)

func main(){
	IsIsogram("alphabe-t")
}

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	fmt.Println(word)

	for i,c := range word { 
		if c== ' ' || c== '-' {
			continue
		}
		
		for j := i + 1; j < len(word); j++ {
			if word[i] == word[j] {
				return false 
			}
		}
	}
	return true
}