package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func sumfromfile()  {
	lines, err := readLines("resultat.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	inp1 := lines[0]
	inp2 := lines[1]
	tall1,_ := strconv.Atoi(inp1)
	tall2,_ := strconv.Atoi(inp2)

	resultat := tall1 + tall2
	if resultat <= 0 {
		fmt.Println("Feil, du må ha et tall som er høyere enn 0")
	}
	f, err := os.OpenFile("resultat.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(f,"\n%d\n", resultat); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}