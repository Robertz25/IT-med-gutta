package main

import (
	"fmt"
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
//-IT-med-gutta , <-- herman, håkon og robert, copyright.
func main() {

	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
