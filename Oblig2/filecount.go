package main

import (

	"fmt"

	"strings"

	"io/ioutil"

	"log"

	"sort"

	"os"
	"bufio"
)

func main() {

	fmt.Println("Information about <filecount.go>: \n")


	file, _ := os.Open("text.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("Number of lines in a file: ", lineCount, "\n")

	fi, err := ioutil.ReadFile("text.txt")

	Stringfile := string(fi)



	if err != nil {

		log.Fatal(err)

	}

	var characterCount = make(map[byte]int)

	for i := 0; i < len(Stringfile); i++ {

		if _, ok := characterCount[Stringfile[i]]; ok {



		} else {

			characterCount[Stringfile[i]] = strings.Count(Stringfile, string(Stringfile[i]))

		}

	}

	type kv struct {

		Key byte

		Value int

	}

	var ss []kv

	for k, v := range characterCount {

		ss = append(ss, kv{k, v})

	}

	sort.Slice(ss, func(i, j int) bool {

		return ss[i].Value > ss[j].Value

	})

	counter := 0

	for _, kv := range ss {

		if counter < 5 {

			fmt.Printf("%d. Rune: %s, counts: %d\n", counter + 1, string(kv.Key), kv.Value)

		}

		counter++

	}

}
