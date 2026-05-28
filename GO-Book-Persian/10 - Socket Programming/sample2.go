package main
import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){

	conn, err := net.Dial("tcp","localhost:5060")

	if err!=nil {
		panic("error establishing connection, make sure the server is up")
	}

	defer conn.Close()

	log.Println("Connected to a tcp server on localhost:5060")
	fmt.Println("Write something to send to server")

	var message string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">> ")
		scanner.Scan()
		message = scanner.Text()
		if message == "exit" {
			return
		}

		fmt.Fprint(conn,message) 
	}
}