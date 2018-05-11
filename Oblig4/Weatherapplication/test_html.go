package main

import (
	"testing"
	"log"
	"os"
	"fmt"
)


func fileFound(filename string) bool {

	if _, err := os.Stat(filename); err != nil {

		if os.IsNotExist(err) {

			return false
		}
	}
	return true
}

func check(e error) {

	if e != nil {

		log.Fatal(e)

	}

}

//tester at templaten Alesund eksisterer.

func testTemplateAlesund(t *testing.T){

	x := fileFound("Alesund.html")

	if x != true {

		t.Error("Template Alesund eksisterer ikke")

	} else{
		fmt.Print("Template Alesund eksisterer")
	}

}

//tester at templaten Tromso eksisterer.

func testTemplateTromso(t *testing.T){

	x := fileFound("Alesund.html")

	if x != true {

		t.Error("Template Tromso eksisterer ikke")

	} else{
		fmt.Print("Templaten Tromso eksisterer")
	}

}

//tester at templaten Stavanger eksisterer.

func testTemplateStavanger(t *testing.T){

	x := fileFound("Stavanger.html")

	if x != true {

		t.Error("Template Stavanger eksisterer ikke")

	} else{
		fmt.Print("Template Stavanger eksisterer")
	}

}
