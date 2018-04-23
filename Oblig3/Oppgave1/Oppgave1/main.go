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

func main() {

	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
