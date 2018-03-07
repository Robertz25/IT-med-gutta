package main

import (
	"fmt"
	"os"
	"log"
)


func main() {
	file, err := os.Open("file.go") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Information about file <" + file + ">:")
}
