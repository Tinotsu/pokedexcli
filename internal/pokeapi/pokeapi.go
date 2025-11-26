// Package pokeapi
package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetNextAreas(offset int) []string {
	strOffset := fmt.Sprintf("%d",offset)
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/?offset=" + strOffset)
	if err != nil {
		fmt.Print("GetNextAreas error: res error\n")
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Print("GetNextAreas error:")
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Print("GetNextAreas error: body error\n")
		log.Fatal(err)
	}

	// var locations []Location
	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Print("GetNextAreas error:")
		log.Fatal(err)
	}
	locationsName := []string{}
	for i := 0; i < len(response.Results); i++ {
		locationsName = append(locationsName, response.Results[i].Name)
	}

	return locationsName
}
