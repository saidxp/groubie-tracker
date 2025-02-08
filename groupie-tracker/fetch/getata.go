package fetch

import (
	"io"
	"log"
	"net/http"
	"encoding/json"
	"groupie-tracker/mod"
)

type Location struct {
    ID            int                 `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"`
}

type GlobalData struct {
    ArtistData mod.Data
    DatesLocations map[string][]string  // Maps locations to arrays of dates
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
		log.Fatal()
	}
	var data Location

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Print(err)
	}
	return data
}

func Mergethedata(artistData *mod.Data, locationData Location) GlobalData {
    locationStruct := locationData
    return GlobalData{
        ArtistData:           *artistData,            
        DatesLocations: locationStruct.DatesLocations,
    }
}
