**3. Forstå prosessadministrajon på et platform**

Evig løkke:
```
package main

import fmt

func main() {
	for {
	fmt.Println("Run Forever");
	}
}
```
Ved å trykke play knappen  eller gå i terminalen og skrive (go run "filnavn".go) blir loopen kjørt over.
Når vi bruker kommandoen ctrl c eller ctrl break klarer vi å avslutte programmet og får da opp meldingen "Process finished with exit code 2."

Vi kan også benytte SIGINT eller SIGQUIT for eksempel, ved å bruke funksjonen Notify (signal.Notify). Kan eventuelt bruke if setning inne i for loopen, til å lage egen måte å stoppe loopen på. 

CPU-en bruker gjennomsnittlig 37.2% av prosessoren, og 733MB minne, som vi leser av i oppgavebehandleren.
