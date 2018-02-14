
**Oppgave 4a**
```
func IterateOverASCIIStringLiteral() {
	for i:= 128; i <= 255; i++{
	fmt.Printf("%X %c %b \n" , i, i, i)
}
}
```
Fra bytes fra 0x80 til 0x9F (128 til 159, desimalverdi) får vi ikke opp symboler. Gruppen bruker windows maskiner der det skjer det samme. Etter dette får vi symboler fra 160 til vi slutter på desimalverdien 255, eller hex FF.

**Oppgave 4b**
```
func main(){
	asciiByteSlice := []byte(asciiNR1)
	fmt.Printf("%s", ExtendedASCIIText(asciiByteSlice))
}

func ExtendedASCIIText(s []byte)string{
	var asciiNR1 = string(s)
	return asciiNR1
}
```
Dersom vi tar å kjører i kommandovinduet:
```
C:\Users\herma\go\src>go run Sorting.go
C:\Users\herma\go\src>€ ÷ ¾ dollar

'€' is not recognized as an internal or external command,
operable program or batch file.

```
Vi får meldingen at den ikke kjenner igjen "€" tegnet.

**Oppgave 4c**
```
package main

import (
	"testing"
	"fmt"
)

package main

import (
	"testing"
	"fmt"
)

func TestExtendedASCIIText(t *testing.T) {
	for i := 0; i<len(nr1AsciiTest);i++{
		if(nr1AsciiTest[i] < 128){
			t.Fail()
			fmt.Println("Value not a part of the extended ascii")
		}
	}
}

C:\Users\herma\go\src>go test
Value not a part of the extended ascii
Value not a part of the extended ascii
Value not a part of the extended ascii
--- FAIL: TestExtendedASCIIText (0.00s)
FAIL
exit status 1

```
Dersom vi kjører go test, i terminal eller kommandovinduet får vi feilmelding, fra func TestExtendedASCIIText(t * testing.T) Årsaken til dette er fordi nr1AsciiTest er mindre enn 128. Extended Ascii table viser oss at det går fra desimale fra 128 til 255. 

```
const nr1AsciiTest = "\x22\x23\x10\x81\x82" <- Dette er test-tall fra opg4_innlevering1.go
```

Hvis vi ser på disse tallene ser vi at 3 av dem er under desimalverdien 128 (x22, x23 og x10). Derfor ser vi litt ovenfor at det blir printet ut 3 ganger at vi har verdier som ikke tilhører extendedAscii.

Dersom vi bruker denne, får vi meldingen pass, fordi alle / verdiene tilhører extended ascii.
```
const PasserTest = "\x90\x95\x97"

func TestExtendedASCIIText(t *testing.T) {
	for i := 0; i<len(PasserTest);i++{
		if(PasserTest[i] < 128){
			t.Fail()
			fmt.Println("Value not a part of the extended ascii")
		}
	}
}

C:\Users\herma\go\src>go test
PASS
ok      _/C_/Users/herma/go/src 0.069s

```
