package main

import (
	"net/http"
	"text/template"
	"log"
	"io/ioutil"
	"fmt"
	"encoding/json"

)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*")) // bruker ParsbGlob fordi vi bruker en mappe
}


func main(){

	http.HandleFunc("/", helloClient) // <- path, opg1
	http.HandleFunc("/1", page1)
	http.HandleFunc("/2", page2)
	http.HandleFunc("/3", page3)
	http.HandleFunc("/4", page4)
	http.HandleFunc("/5", page5)

	http.ListenAndServe(":8080" , nil)
}

//Blir brukt i oppgave1.
func helloClient(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w,"Hello, Client")
}

//Side1 , /1 hentet fra stavanger miljøstasjoner.
func page1(w http.ResponseWriter, r *http.Request)  {

	res, err := http.Get("http://hotell.difi.no/api/json/stavanger/miljostasjoner")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	a := Miljostasjoner{}
	err = json.Unmarshal(body, &a)
	if err != nil {
		panic(err)
	}

	//go template
	tpl.ExecuteTemplate(w, "page1.html", a)

}
// Side2, /2 fylkesnummer.
func page2(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://hotell.difi.no/api/json/difi/geo/kommune?fylke=14")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	b := fylkesnr{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		panic(err)
	}

	//go template
	tpl.ExecuteTemplate(w, "page2.html", b)
}
func page3(w http.ResponseWriter, r *http.Request) {

	res, err := http.Get("https://hotell.difi.no/api/json/valg/valglokaler/2017?")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := kommune{}
	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)

	}

  //go template
	tpl.ExecuteTemplate(w,"page3.html", c)

}
func page4(w http.ResponseWriter, r *http.Request) {

	res, err := http.Get("https://hotell.difi.no/api/json/difi/geo/fylke")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	d := hotell{}
	err = json.Unmarshal(body, &d)
	if err != nil{
		panic(err)

	}


	//go template
	tpl.ExecuteTemplate(w,"page4.html", d)
}
func page5(w http.ResponseWriter, r *http.Request) {

	res, err := http.Get("https://hotell.difi.no/api/json/brreg/enhetsregisteret?page=8")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	e := selskap{}
	err = json.Unmarshal(body, &e)
	if err != nil {
		panic(err)

	}
  
  //go template
	tpl.ExecuteTemplate(w,"page5.html", e)

}


//struct for func page1.
type Miljostasjoner struct {
	Entries  []struct {
		Latitude     string `json:"latitude"`
		Navn         string `json:"navn"`
		Plast        string `json:"plast"`
		Glass_metall string `json:"glass_metall"`
		Tekstil_sko  string `json:"tekstil_sko"`
		Longitude    string `json:"longitude"`
	} `json:"entries"`

	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

//struct for func page2.
type fylkesnr struct {
	Entries []struct {
		Kommune string `json:"kommune"`
		Fylke   string `json:"fylke"`
		Navn    string `json:"navn"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

//struct for func page3.
type kommune struct {
	Entries []struct {
		Area              string `json:"area"`
		BoroughName       string `json:"borough_name"`
		AddressLine       string `json:"address_line"`
		PollingPlaceID    string `json:"polling_place_id"`
		BoroughID         string `json:"borough_id"`
		CountyName        string `json:"county_name"`
		PollingPlaceName  string `json:"polling_place_name"`
		GpsCoordinates    string `json:"gps_coordinates"`
		MunicipalityID    string `json:"municipality_id"`
		CountyID          string `json:"county_id"`
		MunicipalityName  string `json:"municipality_name"`
		ElectionDayVoting string `json:"election_day_voting"`
		InfoText          string `json:"info_text"`
		OpeningHours      string `json:"opening_hours"`
		PostalCode        string `json:"postal_code"`

	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`

}

//struct for func page4.
type hotell struct {
	Entries []struct {
		Navn   string `json:"navn"`
		Nummer string `json:"nummer"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

//struct for func page5.
type selskap struct {
	Entries []struct {
		Tvangsavvikling   string `json:"tvangsavvikling"`
		Regnskap          string `json:"regnskap"`
		Forradrpostnr     string `json:"forradrpostnr"`
		AnsatteAntall     string `json:"ansatte_antall"`
		Postadresse       string `json:"postadresse"`
		Nkode3            string `json:"nkode3"`
		Ppoststed         string `json:"ppoststed"`
		Konkurs           string `json:"konkurs"`
		Stiftelsesdato    string `json:"stiftelsesdato"`
		Sektorkode        string `json:"sektorkode"`
		AnsatteDato       string `json:"ansatte_dato"`
		Organisasjonsform string `json:"organisasjonsform"`
		Navn              string `json:"navn"`
		Regifriv          string `json:"regifriv"`
		Forradrkommnr     string `json:"forradrkommnr"`
		Regimva           string `json:"regimva"`
		TlfMobil          string `json:"tlf_mobil"`
		Forradrland       string `json:"forradrland"`
		Ppostland         string `json:"ppostland"`
		Avvikling         string `json:"avvikling"`
		Regifr            string `json:"regifr"`
		Hovedenhet        string `json:"hovedenhet"`
		Forretningsadr    string `json:"forretningsadr"`
		URL               string `json:"url"`
		Forradrpoststed   string `json:"forradrpoststed"`
		Tlf               string `json:"tlf"`
		Nkode1            string `json:"nkode1"`
		Nkode2            string `json:"nkode2"`
		Forradrkommnavn   string `json:"forradrkommnavn"`
		Regdato           string `json:"regdato"`
		Orgnr             string `json:"orgnr"`
		Regiaa            string `json:"regiaa"`
		Ppostnr           string `json:"ppostnr"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}
