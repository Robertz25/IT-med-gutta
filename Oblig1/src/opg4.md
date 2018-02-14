
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
package main

import "fmt"

func ExtendedAsciiText(x string) string {
	r := make([]rune, len(x))
	for i := 0; i < len(x); i++ {
		r[i] = cp1252[x[i]]
	}
	return string(r)
}


func main() {
	x := "\x80\xf7\xbe dollar"
	str := ExtendedAsciiText(x)
	fmt.Printf("%q\n", str)
}

func init() {
	for i, r := range cp1252 {
		if r == 0 {
			cp1252[i] = rune(i)
		}
	}
}

var cp1252 = [255]rune {
	0x80: '\u20AC', // EURO SIGN
	0x81: '\uFFFD', // UNDEFINED
	0x82: '\u201A', // SINGLE LOW-9 QUOTATION MARK
	0x83: '\u0192', // LATIN SMALL LETTER F WITH HOOK
	0x84: '\u201E', // DOUBLE LOW-9 QUOTATION MARK
	0x85: '\u2026', // HORIZONTAL ELLIPSIS
	0x86: '\u2020', // DAGGER
	0x87: '\u2021', // DOUBLE DAGGER
	0x88: '\u02C6', // MODIFIER LETTER CIRCUMFLEX ACCENT
	0x89: '\u2030', // PER MILE SIGN
	0x8A: '\u0160', // LATIN CAPITAL LETTER S WITH CARON
	0x8B: '\u2039', // SINGLE LEFT-POINTING ANGLE QUOTATION MARK
	0x8C: '\u0152', // LATIN CAPITAL LIGATURE OE
	0x8D: '\uFFFD', // UNDEFINED
	0x8E: '\u017D', // LATIN CAPITAL LETTER Z WITH CARON
	0x8F: '\uFFFD', // UNDEFINED
	0x90: '\uFFFD', // UNDEFINED
	0x91: '\u2018', // LEFT SINGLE QUOTATION MARK
	0x92: '\u2019', // RIGHT SINGLE QUOTATION MARK
	0x93: '\u201C', // LEFT DOUBLE QUOTATION MARK
	0x94: '\u201D', // RIGHT DOUBLE QUOTATION MARK
	0x95: '\u2022', // BULLET
	0x96: '\u2013', // EN DASH
	0x97: '\u2014', // EM DASH
	0x98: '\u02DC', // SMALL TILDE
	0x99: '\u2122', // TRADE MARK SIGN
	0x9A: '\u0161', // LATIN SMALL LETTER S WITH CARON
	0x9B: '\u203A', // SINGLE RIGHT-POINTING ANGLE QUOTATION MARK
	0x9C: '\u0153', // LATIN SMALL LIGATURE OE
	0x9D: '\uFFFD', // UNDEFINED
	0x9E: '\u017E', // LATIN SMALL LETTER Z WITH CARON
	0x9F: '\u0178', // LATIN CAPITAL LETTER Y WITH DIAERESIS
}

C:\Users\herma\go\src>go run test.go
"€÷¾ dollar"

```
Som vi ser får vi skrevet ut €÷¾ dollar. Som vi ser under var cp1252 = [255]rune, er det listet mange forskjellige tegn fra extended ASCII, vi behøver kun den øverste (0x80) egentlig. Disse tre tegnene: €÷¾ tilhører extended ASCII, men ÷¾ kan vi fortsatt bruke siden de er matematiske symboler.

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
