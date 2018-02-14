package main

import (
	"testing"
	"fmt"
)


func TestExtendedASCIIText(t *testing.T) {
	for i := 0; i<len(nr1AsciiTest);i++{
		if(nr1AsciiTest[i]) < 128{
			t.Fail()
			fmt.Println("Value not a part of the extended ascii")
		}
	}
}
