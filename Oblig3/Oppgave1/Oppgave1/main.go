package main

import (
	"fmt"
	"net"
	"os"
	"net/http"
	"log"
)

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}


    func HelloServer(w http.ResponseWriter, req *http.Request) {
        fmt.Fprint(w, "Hello, client\n")
  }

func main() {
	ServerAddr,err := net.ResolveUDPAddr("udp",":8080")
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

   http.HandleFunc("/", HelloServer)
    log.Fatal(http.ListenAndServe(":8080", nil))



	for {
		n,addr,err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ",string(buf[0:n]), " from ",addr)

		if err != nil {
			fmt.Println("Error: ",err)
		}
		ServerConn.WriteToUDP([]byte("ack"), addr)
	}
}
