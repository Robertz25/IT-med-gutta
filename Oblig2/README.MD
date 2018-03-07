<h1> IS105, obligatorisk oppgave 2 </h1>

**Oppgave 1A**
```
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
```

**Oppgave 2**

```
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
//	fmt.Println("Number of lines in a file:", lineCount)

//	for  i := 1; i < 6; i++{
//			fmt.Println(i,".Rune:", i ," Counts:", lineCount)
		}

	}
```
