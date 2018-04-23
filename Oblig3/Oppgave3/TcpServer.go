package main

import (
	"net"
	"log"
	"io"
	"fmt"
)

func main() {
li, err := net.Listen("tcp", ":8080")
if err != nil {
	log.Panic(err)
}
defer li.Close()

for {
	conn, err := li.Accept()
	if err != nil {
		log.Println(err)
	}
	io.WriteString(conn, "\n Hello from TCP server \n")
	fmt.Fprintln(conn, "How is your day")
	fmt.Fprintf(conn, "%v", "Well I hope!")

	conn.Close()
}

}
