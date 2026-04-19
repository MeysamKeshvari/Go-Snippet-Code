package main 

import (
	"fmt"
	"net"
)

func main(){

	listener,err := net.Listen("tcp","localhost:5060")

	if err!=nil{
		panic("can not create a listener , error:" + err.Error())
	}

	defer listener.Close()
	fmt.Println("Server started, ready to receive connection")

	var i = 1 
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("error in accepting connection , error:" + err.Error())
		}

		go handleConnection(conn, i)
		i++
	}
}

func handleConnection(conn net.Conn, i int){
	fmt.Println("A new connection Established number assigned: " , i)
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n,err := conn.Read(buffer)
		if err!= nil{
			fmt.Println("error in reading the connection , error:" + err.Error())
			return
		}
		fmt.Printf("Client %d: %s\n" , i ,string(buffer[:n]))
	}
}