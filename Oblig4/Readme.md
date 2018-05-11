**Herman Fensbekk - hermanfensbekk@hotmail.com**

**Robert Zakariassen - robert.zakariassen@lyse.net**

**Håkon Sveen - hakon_s_s@hotmail.com**

# Obligatorisk oppgave 4 #


## Systemspesifikasjon ##

På obligatorisk oppgave 4 har vi valgt å lage et program som sammenligner temperaturer fra forskjellige byer i Norge.
Byene vi har valgt å sammenligne er Ålesund, Tromsø og Stavanger. Vi tenkte det var greie byer å sammenligne for å få en oversikt over litt forskjellige steder i Norge. For å få tak i API til de ulike stedene brukte vi openweathermap.org. Der måtte vi lage en bruker for å få en egen API key som kunne brukes til å hente API for forskjellige steder. For å få en litt mer visuell oversikt har vi brukt en widget fra Weatherwidget.io. Vi har også brukt switch til å legge inn anbefalinger til hva klær som passer til de ulike gradene eller fortelle om hvordan været er. Noen eksempler er hvis det er over 17 grader og mindre enn 4m/s i vindstyrke kommer det opp en melding om at det er "Supert sommervær og lite vind, finn fram badetøy!". Hvis det fortsatt er 17 grader eller mer, men over 4m/s i vindstyrke kommer det opp en melding tilsvarende "Flott sommervær, men en del vind idag!". Hvis det er under 4 grader kommer det opp en melding om at det er kaldt ute og at det kan være lurt å kle på seg godt.

## Systemarkitektur ##

Applikasjonen vår består av en webserver som henter API fra Openweathermap og gjør dataen om til lettleselig tekst som er enkelt for brukeren å forstå. Stedene vi henter data fra er Tromsø, Ålesund og Stavanger.

Rå API for Tromsø:

![Alt text](https://github.com/Robertz25/IT-med-gutta/blob/master/Oblig4/Bilder/jsonApiTromso.png?raw=true)

API for Tromsø etter konvertering:

![Alt text](https://github.com/Robertz25/IT-med-gutta/blob/master/Oblig4/Bilder/API%20konvertert%20.png?raw=true)

Vi har også lagt inn en widget fra Weatherwidget.io og bilder fra de forskjellige stedene for å gjøre nettsiden estetisk bedre å se på og lettere å forstå.

**Link til oppgaven**

[Link til applikasjonen](https://github.com/Robertz25/IT-med-gutta/blob/master/Oblig4/Weatherapplication/main.go)
[Link til folderen](https://github.com/Robertz25/IT-med-gutta/tree/master/Oblig4/Weatherapplication)
