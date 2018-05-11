package main

import (
	"net/http"
	"log"
	"encoding/json"
	"testing"
	"fmt"
)

// Henter ut Apiet fra openweathermap -> Alesund.
func AlesundApi(url string) Side1 {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=%C3%85lesund&APPID=12576bc04cecb325c71bbe3972592fa1&units=metric")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var data Side1
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data

}

//Tester at navnet på byen er Ålesund.
func testAlesund(t *testing.T) {
	x := AlesundApi("http://api.openweathermap.org/data/2.5/weather?q=%C3%85lesund&APPID=12576bc04cecb325c71bbe3972592fa1&units=metric")
	if x.Name != "Ålesund" {
  
		// tester hvis navnet på byen,  Ålesund er rett. Kjører en error hvis ikke, og hvilket navn som ble returnert istedet.
    
		t.Errorf("Bynavnet som ble returnert var ikke riktig.", x.Name)
	} else{
		fmt.Print("OpenWeatherMap returnerer Ålesund, som er riktig.")
	}
}

// Henter ut Apiet fra openweathermap -> Tromso.
func TromsoApi(url string) Side2 {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Tromso&APPID=12576bc04cecb325c71bbe3972592fa1&units=metric")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var data Side2
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data

}

//Tester at navnet på byen er Tromso.
func testTromso(t *testing.T) {
	x := TromsoApi("http://api.openweathermap.org/data/2.5/weather?q=Tromso&APPID=12576bc04cecb325c71bbe3972592fa1&units=metric")
	if x.Name != "Tromso" {
  
		// tester hvis navnet på byen, Tromso er rett. Kjører en error hvis ikke, og hvilket navn som ble returnert istedet.
    
		t.Errorf("Bynavnet som ble returnert var ikke riktig.", x.Name)
	} else{
		fmt.Print("OpenWeatherMap returnerer Tromso, som er riktig.")
	}
}

// Henter ut Apiet fra openweathermap -> Stavanger.
func StavangerApi(url string) Side3 {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Stavanger&APPID=12576bc04cecb325c71bbe3972592fa1&units=metric")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var data Side3
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data

}

//Tester at navnet på byen er Stavanger.
func testStavanger(t *testing.T) {
	x := TromsoApi("http://api.openweathermap.org/data/2.5/weather?q=Stavanger&APPID=12576bc04cecb325c71bbe3972592fa1&units=metric")
	if x.Name != "Stavanger" {
  
		// tester hvis navnet på byen, altså Stavanger er rett. Kjører en error hvis ikke, og hvilket navn som ble returnert istedet.
    
		t.Errorf("Bynavnet som ble returnert var ikke riktig.", x.Name)
	} else{
		fmt.Print("OpenWeatherMap returnerer Stavanger, som er riktig.")
	}
}
