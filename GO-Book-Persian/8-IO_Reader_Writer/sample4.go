package main 

import (
	"io"
	"os"
)
func SayHello(w io.Writer){
	w.Write([]byte("Hello World"))
}

func main(){
	file,_ := os.Create("hello.txt")
	defer file.Close()
	SayHello("file")
} 