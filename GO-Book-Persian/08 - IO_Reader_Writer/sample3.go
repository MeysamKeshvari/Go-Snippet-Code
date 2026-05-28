package main 

import(
	"io" 
	"os"
)

func SayHello(w io.Writer){
	w.Write([]byte("Hello world"))
}

func main(){
	SayHello(os.Stdout)
}

