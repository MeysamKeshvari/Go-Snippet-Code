package main 

import "fmt"

type Str string

func (s *Str) Write(b []byte) (int, error){
	*s += Str(string(b))
	return len(b),nil
}

func main(){
	var str Str 
	str.Write([]byte("Hello World"))
	fmt.Println(str)
} 