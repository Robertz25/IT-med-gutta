package main

import (
	"net"
	"log"
	"fmt"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil{
		log.Fatalln(err)
	}

	_, err = conn.Write([]byte("Hello server! - from client "))
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println("Message sent: Hello server! - from client ")

	_, err = conn.Write([]byte("This is a random message"))
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println("Message sent: This is a random message")

	for{
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil{
			fmt.Println("Connection closed")
			return
		}
		data := buffer[:dataSize]
		fmt.Println("Received message:", string(data))
	}
}

