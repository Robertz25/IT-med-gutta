package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("text.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("Number of lines in a file:", lineCount)

	for  i := 1; i < 6; i++{
			fmt.Println(i,".Rune:", i ," Counts:", lineCount)
		}

	}
