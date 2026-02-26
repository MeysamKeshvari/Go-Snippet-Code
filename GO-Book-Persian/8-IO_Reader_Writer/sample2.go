package main 

import "io"
import "fmt"

type String string

func (s *String) Write(b []byte) (int, error){
	*s += String(string(b))
	return len(b), nil
}  

func SayHello(w io.Writer){
	w.Write([]byte("Hello World"))
}
func main(){
	var str String 
	SayHello(&str)
	fmt.Println(str)
}

