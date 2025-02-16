package fetch

import (
	"io"
	"log"
	"net/http"
	"encoding/json"
	"groupie-tracker/mod"
)

type Location struct {
    ID            int                `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"` //TODO: MAZAL KHAS N7AYAD CHI HAJA
}

type GlobalData struct {
    ArtistData mod.Data
    DatesLocations map[string][]string  
}

type City struct {
	ID int	`json:"id"`
	Loca []string `json:"locations"`
}

type Year struct {
	ID int `json:"id"`
	City []string `json:"dates"`
}

func Locationanddate(s string) Location {
	// her in location there is the date to of the location where is 
	res , err  := http.Get(s)

	if err != nil {
		log.Fatal("Errot")
	}
	defer res.Body.Close()

	body , err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error")
	}
	var data Location

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Print(err) 
	}
	return data
}
func Datesof(s string) Year {
	r , err := http.Get(s)
	if err != nil {
		log.Fatal()
	}
	body , err  := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal()
	}
	var data Year

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal()
	}
	return data
}

func Locationof(s string) City {
	 
	response , err := http.Get(s)

	if err != nil {
		log.Fatal("Erore fetch data")
	}
	body , err := io.ReadAll(response.Body)
	
	if err != nil {
		log.Fatal()
	}
	var Data City
	err = json.Unmarshal(body, &Data)
	if err != nil {
		log.Fatal()
	}	
	return Data
}

func Mergethedata(artistData *mod.Data, locationData Location) GlobalData {
    locationStruct := locationData
    return GlobalData{
        ArtistData:           *artistData,            
        DatesLocations: locationStruct.DatesLocations,
    }
}
