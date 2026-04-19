package main 

import (
	"fmt"
	"io"
)
type String struct {
	data string 
	pos int
}

func (s *String) Read(b []byte) (int,error) {
	n:= copy(b, s.data[s.pos:])
	s.pos += n

	var err error 
	if s.pos >= len(s.data){
		err = io.EOF
	}
	return n, err
}

func (s *String) Write(b []byte) (int, error){
	s.data += string(b)
	return len(b), nil
}

func main(){
	var str String 
	str.Write([]byte("Hello World"))
	buffer := make([]byte, 20)
	n, err:= str.Read(buffer)

	fmt.Println("Bytes read:" ,n)
	fmt.Println("Error:",err)
	fmt.Println("Buffer:", string(buffer))
	
}
