| **Binære tall (mest signifikant bit til venstre**  |    **Hexadesimaltall**  |  **Desimaltall**|
|------------------------------------------------|---------------------|-------------|
| 1101 |  0xD| 13 |
| 1101 1110 1010 |    DEA  |   3562 |
|1010 1111 0011 0100| 0xAF34 |    44652 |
| 1111 1111 1111 1111| FFFF |   65535 |
| 1000 1011 1100 01010| 1178A |    71562 |

**Oppgave 1a**

**Binære tall til hexadesimale tall:**

Vi gjør fire og fire siffer i det binære tallet om til et hexadesimalt siffer. Vi begynner bakerst i det binære tallet, til høyre og velger de fire bakerste sifrene. Deretter plukker vi fire nye siffer, helt til det ikke er flere igjen. Dersom det er færre enn fire siffer igjen må vi fylle ut de manglende sifrene med 0.
Vi bruker 100011. Dersom vi begynner bakerst og deler opp får vi: 0011. de resterende er 10, vi legger bare 00 foran og får da: 0010 0011.

|**Binært tall** |**Hexadesimalt tall**|
|------------|-------------------|
|0011 |	3 |
|0010 |	2 |

Leser nedenfra og opp, og får det hexadesimale tallet 23.

**Hexadesimale tall til binære tall:**

Den enkleste måten å gå fra hexadesimalt tall til et binært tall er ved å se på hvert enkelt siffer i det hexadesimale tallet og finne det tilsvarende binære tallet, for så å slå de binære tallene vi har funnet sammen i akkurat den samme rekkefølgen som de sto i det hexadesimale tallet.

|Heksadesimalt tall	| Binært tall|
|-------------------|------------|
|2 |0010|
|3 |0011|

**Binære tall til desimale tall:**

For å gå fra binære tall til desimale tall, må man gange sifferet (som består av kun sifrene 0 og 1) med grunntallet 2 opphøyet i sifferets posisjon. Deretter må man addere med neste tall i rekken. Det kan se slik ut. 101 -> (1*2^2) + (0*2^1) + (1*2^0). Vi legger sammen: 4 + 0 + 1 = 5. Som vi ser i det binære tallet 101, er det tre posisjoner. Årsaken til at tallet helt til venstre er opphøyet i 2 i stedet for 3, er fordi sifferet helt til høyre står i tallrekke 0.
Om x = 1 eller 0, og 0-te desimal posisjon er lik n, kan regningen skrives slik: …(x * 2^n+1) + (x * 2^n)

Binært tall (101001), omgjøres til desimal slik, 1*2^5 + 0*2^4 + 1*2^3 + 0*2^2 + 0*2^1 + 1*2^0 =    32 + 0 + 8 + 0 + 0 + 1 = 41.

**Desimale tall til binære tall:**

For å gå fra desimale til binære tall gjør vi slik, at vi deler tallet på to, dersom vi får rest, da skriver vi ned 1, hvis vi ikke får rest skriver vi 0. Vi benytter heltallsdivisjon, vi benytter ikke desimaler og runder heller ned (for eksempel 40.5 rundes ned til 40). Det er også slik at tallet etter divisjon er oddetall, står vi igjen med rest, og motsatt ved partall, %modulus regning. Vi bruker heltallsdivisjon helt til vi står igjen med 0. Når vi er ferdige leses rest sifrene fra bunnen og oppover.
Bruker eksempelet ovenfor, med desimale 41.

|verdi	|utregning	|rest|
|-------|-----------|----|
|41|	41/2 = 20, 41%2 = 1|	1|
|20|	20/2 = 10, 20%2 = 0|	0|
|10|	10/2 = 5, 10%2 = 0|	0|
|5|	5/2 = 2, 5%2 = 1|	1|
|2|	2/2 = 1, 2%2 = 0|	0|
|1|	1/2 = 0, 1%2 = 1|	1|

Desimale 41, leses av fra sifrene rest fra bunnen og oppover. Da får vi det binære tallet 101001.

**Oppgave 1b**

**Hexadesimale tall til desimale tall:**
Det hexadesimale tallsystemet har grunntallet 16, og har verdiene 0 til 9, og a til og med f. Der a er 10, til f som er 15. Vi prøver å regne hexadesimale ABC5 til desimale:

|Posisjon	|Siffer|	Utregning|
|---------|------|-----------|
|0|5|	5*16^0 = 5|
|1|C|	12*16^1 = 192|
|2|B|	11*16^2 = 2816|
|3|A|10*16^3 = 40960|

Hexadesimale ABC5 er desimal tallet:  5 + 192 + 2816 + 40960 = 43973.

**Desimale tall til hexadesimale tall:**

Utregningen her blir nesten samme som fra desimal til binær tall, men vi deler på 16 i stedet for 2.

|Verdi|	Utregning	Hexadesimal| rest|
|-----|----------------------|-----|
|43973|	43973/16 = 2748, 43973%16|	5|
|2748|2748/16 = 171, 2748%16|	12 = C|
|171|	171/16 = 10, 171%16|11 = B|
|10|	10/16 = 10%16|10 = A|

Desimal tallet 43973 er (lest fra bunnen og opp) det hexadesimale tallet ABC5. 



