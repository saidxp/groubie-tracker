package main

import (

	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"groupie-tracker/fetch"
	"groupie-tracker/mod"

)


func get() any {

	//the first i need to fetch the data from api
	url := "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(url)
	if err != nil {
    	log.Fatal(err)
	}
	defer resp.Body.Close() 

	fmt.Println("Status Code:", resp.Status) 

	body, err := io.ReadAll(resp.Body) // Read response body
	if err != nil {
    	log.Fatal(err)	
	}
	var slice []mod.Data
	err = json.Unmarshal(body, &slice)

	return slice 
}

func root(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	data := get().([]mod.Data)
	tmp := template.Must(template.ParseFiles("tamplate/hello.html"))
	tmp.Execute(w,data)
}

func artist(w http.ResponseWriter, r *http.Request) {
	//i will fetch the id of the id who is the client it's want !!
	targerid := r.URL.Query().Get("id")
	//then i will do match the id of the who it's come from client and the the who i fetch from API !!
	thedataofApi := get().([]mod.Data)
	number, err := strconv.ParseInt(targerid, 10, 64)
	if err != nil {
    	http.Error(w, "Invalid ID", http.StatusBadRequest)
    	return
	}
	//Then i will take the data struct who is match whith the target id :
	var Str mod.Data	
	for _, art := range thedataofApi {
		if art.ID == int(number) {
			Str = art
		}
	}
	/*
	"locations": "https://groupietrackers.herokuapp.com/api/locations/1",
    "concertDates": "https://groupietrackers.herokuapp.com/api/dates/1",
    "relations": "https://groupietrackers.herokuapp.com/api/relation/1"
	*/	
	Struct1 := fetch.Locationanddate(Str.Relations)
	Global := fetch.Mergethedata(&Str,Struct1)
	tmp := template.Must(template.ParseFiles("tamplate/artist.html"))
	tmp.Execute(w,Global)


}

func main() {

	http.HandleFunc("/", root)
	http.HandleFunc("/artist", artist)
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080",nil) 

}