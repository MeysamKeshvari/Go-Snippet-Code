package main 

import (
	"fmt"
	"io"
)


type String struct {
	data string 
	pos int
}

func main(){
	var str String 
	io.WriteString(&str,"this is a text without casting to []byte")
	fmt.Println(str.data)
}


