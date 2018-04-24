package main

import (
	"fmt"
	"net"
	"log"
)

func main(){
	fmt.Println("Server listening on 8080")
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil{
		log.Fatalln(err)
	}
	defer listener.Close()

	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Fatalln(err)
		}
		fmt.Println("Server and client connected")

		go listenConnection(conn)
	}
}
func listenConnection(conn net.Conn){
	for{
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil{
			fmt.Println("Connection closed.")
			return
		}
		data := buffer[:dataSize]
		fmt.Println("Recieved message: ", string(data))

		_,err = conn.Write(data)
		if err!= nil{
			log.Fatalln(err)
		}
		fmt.Println("Message sent: " ,string(data))
	}

}


