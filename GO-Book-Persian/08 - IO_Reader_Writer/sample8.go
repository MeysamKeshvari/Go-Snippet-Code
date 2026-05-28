package main 

import (
	"fmt"
	"io"
)

type String struct {
	data string 
	pos int
}

func ReadWriteSomething(rw io.ReadWriter) string {
	rw.Write([]byte("Go Language is fun!"))
	buff := make([]byte,50)
	n,_ := rw.Read(buff)
	return string(buff[:n])  
}
func main(){
	var str String
	result := ReadWriteSomething(&str)
	fmt.Println(result)
}


