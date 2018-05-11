package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"github.com/parnurzeal/gorequest"
	"fmt"
	"errors"
)
func main() {
	http.HandleFunc("/", Alesund)
	http.HandleFunc("/1", Tromso)
	http.HandleFunc("/2", Stavanger)
	http.HandleFunc("/search/", handlerSearch)
	http.HandleFunc("/home/", handler)
	http.HandleFunc("/page1/", handlerNav)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}


}
func handler(w http.ResponseWriter, r *http.Request) {
	d, err := getData("Oslo", w, r)
	if err != nil {
		return
	}
	render(w, d, "template1.html")
}
func Alesund (w http.ResponseWriter, r *http.Request) {

	//&units=metric. Henter api ut i det metriske systemet, istedet for imperial.

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=%C3%85lesund&APPID=12576bc04cecb325c71bbe3972592fa1&units=metric")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	jsonErr := json.Unmarshal(body, &verdi1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	MeldingAlesund()

	temp, err := template.ParseFiles("Alesund.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, verdi1)
	if err != nil {
		log.Fatal(err)
	}

}
func Tromso (w http.ResponseWriter, r *http.Request) {

	//&units=metric. Henter api ut i det metriske systemet, istedet for imperial.

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Tromso&APPID=12576bc04cecb325c71bbe3972592fa1&units=metric")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &verdi2)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	MeldingTromso()

	temp, err := template.ParseFiles("Tromso.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, verdi2)
	if err != nil {
		log.Fatal(err)
	}

}
func Stavanger (w http.ResponseWriter, r *http.Request) {

	//&units=metric. Henter api ut i det metriske systemet, istedet for imperial.

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Stavanger&APPID=12576bc04cecb325c71bbe3972592fa1&units=metric")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &verdi3)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	MeldingStavanger()

	temp, err := template.ParseFiles("Stavanger.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, verdi3)
	if err != nil {
		log.Fatal(err)
	}
}
func MeldingAlesund () {

	//Slipper å lage variabel for grader i celsius, der vi trekker fra fahrenheit, siden vi brukte units=metric, da vi hentet apien.

	switch {
	case verdi1.Main.Temp >= 18 && verdi1.Wind.Speed < 4:
		verdi1.Main.TempA = "Supert sommervær og lite vind, finn fram badetøy!"
	case verdi1.Main.Temp >= 18 && verdi1.Wind.Speed >= 4:
		verdi1.Main.TempA = "Flott sommervær, men en del vind idag!"
	case verdi1.Main.Temp <= 18 && verdi1.Main.Temp > 10:
		verdi1.Main.TempA = "Grei temperatur ute."
	case verdi1.Main.Temp <= 10 && verdi1.Main.Temp >= 4:
		verdi1.Main.TempA = "Litt kjølig idag, kle på deg godt."
	case verdi1.Main.Temp < 4:
		verdi1.Main.TempA = "Brrr... kaldt idag, hold deg innendørs eller kle  deg smart."
	}
}
func MeldingTromso () {

	//Slipper å lage variabel for grader i celsius, der vi trekker fra fahrenheit, siden vi brukte units=metric, da vi hentet apien.

	switch {
	case verdi2.Main.Temp >= 18 && verdi2.Wind.Speed < 4:
		verdi2.Main.TempB = "Flott nordnorsk sommervær med lite vind, kom deg ut!"
	case verdi2.Main.Temp >= 18 && verdi2.Wind.Speed >= 4:
		verdi2.Main.TempB = "Bra nordnorsk sommervær, men en del vind idag!"
	case verdi2.Main.Temp <= 18 && verdi2.Main.Temp > 10:
		verdi2.Main.TempB = "Helt grei temperatur ute, ingen vits i å sitte inne."
	case verdi2.Main.Temp <= 10 && verdi2.Main.Temp >= 4:
		verdi2.Main.TempB = "Ikke spesielt varmt, eller kaldt."
	case verdi2.Main.Temp < 4:
		verdi2.Main.TempB = "Godt og varmt, dersom du er nordlending ;)"
	}
}
func MeldingStavanger () {

	//Slipper å lage variabel for grader i celsius, der vi trekker fra fahrenheit, siden vi brukte units=metric, da vi hentet apien.

	switch {
	case verdi3.Main.Temp >= 18 && verdi3.Wind.Speed < 4:
		verdi3.Main.TempC = "Kanon vær idag, ingen vind!"
	case verdi3.Main.Temp >= 18 && verdi3.Wind.Speed >= 4:
		verdi3.Main.TempC = "Kanon vær, men en del vind!"
	case verdi3.Main.Temp <= 18 && verdi3.Main.Temp > 10:
		verdi3.Main.TempC = "Helt ok temperatur ute"
	case verdi3.Main.Temp <= 10 && verdi3.Main.Temp >= 4:
		verdi3.Main.TempC = "Litt kaldt idag."
	case verdi3.Main.Temp < 4:
		verdi3.Main.TempC = "Veldig kaldt idag, bli inne eller kle på deg."
	}
}

type Side1 struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		Description int  `json:"description"`
		Sys        Sys       `json:"sys"`


		TempA    string
	} `json:"main"`
	Wind       struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	ID int `json:"id"`
	Name string `json:"name"`


}

type Side2 struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempB    string
	} `json:"main"`
	weather struct {
		Description int  `json:"description"`
	}
	Wind       struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	ID int `json:"id"`
	Name string `json:"name"`
}
type Side3 struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		Description int  `json:"description"`
		TempC    string
	} `json:"main"`
	Wind       struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	ID int `json:"id"`
	Name string `json:"name"`


}


var verdi1 Side1
var verdi2 Side2
var verdi3 Side3




func getData(location string, w http.ResponseWriter, r *http.Request) (Side1, error) {
	var d Side1
	requestP1 := "http://api.openweathermap.org/data/2.5/weather?q="
	requestP2 := "&APPID=12576bc04cecb325c71bbe3972592fa1&units=metric"
	_, body, err := gorequest.New().Get(requestP1 + location + requestP2).End()
	if err != nil {
		log.Fatal("Error in request")
	}
	// fmt.Println("It is :", string(body))
	errs := json.Unmarshal([]byte(body), &d)
	if errs != nil {
		render(w, d, "NotFound.html")
		fmt.Fprintf(w, "<h1>Error in getting request from API</h1>")
		return d, errors.New("Wrong request")
	}
	return d, nil
}

func handlerNav(w http.ResponseWriter, r *http.Request) {
	d, err := getData(r.URL.Path[len("/search/"):], w, r)
	if err != nil {
		return
	}
	render(w, d, "template1.html")
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	d, err := getData(r.URL.Path[len("/search/"):], w, r)
	if err != nil {
		return
	}
	render(w, d, "template1.html")
}
func render(w http.ResponseWriter, d Side1, filename string) {
	t, errr := template.ParseFiles(filename)
	if errr != nil {
		log.Fatal("Error in parsing tempalte")
	}
	t.Execute(w, d)
}
