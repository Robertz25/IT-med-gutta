package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Printf("couldn't open file: %s\n", err)
		os.Exit(1)
	}

	//Get stats
	stats, err := file.Stat()
	if err != nil {
		fmt.Printf("couldn't get stats: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Size: %d bytes\n", stats.Size())
	fmt.Printf("Is dir: %t\n", stats.IsDir())
	fmt.Printf("Is regular: %t\n", stats.Mode().IsRegular())
	fmt.Printf("Unix permissions: %s\n", stats.Mode().String())
	fmt.Printf("Append only: %t\n", stats.Mode()&os.ModeAppend != 0)
	fmt.Printf("Is device: %t\n", stats.Mode()&os.ModeDevice != 0)
	fmt.Printf("Is symlink: %t\n", stats.Mode()&os.ModeSymlink != 0)

}
