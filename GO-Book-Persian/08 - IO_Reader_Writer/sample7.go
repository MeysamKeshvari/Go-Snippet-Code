package main 

import (
	"fmt"
	"os"
)

func main(){
	buffer := make([]byte,10)
	n,_:= os.Stdin.Read(buffer)
	fmt.Println("buffer:", string(buffer[:n]))
}

