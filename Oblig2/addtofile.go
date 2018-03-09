package main
import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os/signal"
	"syscall"
)

func main() {
	addtofile()
	sumfromfile()
	readResult("resultat.txt")
}

func addtofile() {

	var tall1 int
	var tall2 int

	//I terminal så skriver brukeren inn 2 tall, nummer 1 & nummer 2
	fmt.Println("Skriv inn et tall: ")
	fmt.Scan(&tall1)
	if tall2 <= 0 {
		fmt.Println("Error, du må skrive inn et tall, eller et tall som er høyere enn 0")
	}
	fmt.Println("Skriv inn et tall: ")
	fmt.Scan(&tall2)
	if tall1 <= 0 {
		fmt.Println("Feil, du må skrive inn et tall som er større enn 0 ")
	}
	file, err := os.Create("resultat.txt")
	if err != nil {
		log.Fatal("Det er noe feil, kan ikke lage filen", err)
	}
	defer file.Close()

	f, err := os.OpenFile("resultat.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(f, "%d\n%d", tall1, tall2); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readResult(path string) {
	sign :=make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT)

	data, err := ioutil.ReadFile(path)
	checkErr(err)
	tempData := string(data)
	stringData := strings.Split(tempData, "\n")
	temp := stringData[len(stringData)-2]
	resultat, err := strconv.Atoi(temp)
	checkErr(err)

	fmt.Println("Dette er resultatet: ", resultat)
	run:=<-sign
	fmt.Print(run)
}