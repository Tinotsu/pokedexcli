// Package pokeapi
package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	"github.com/Tinotsu/pokedexcli/internal/pokecache"
	"time"
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
	interval := time.Second * 30
	strOffset := fmt.Sprintf("%d",offset)
	url := "https://pokeapi.co/api/v2/location-area/?offset=" + strOffset
	cache := pokecache.NewCache(interval)
	content, isCached  := cache.Get(url)
	if isCached {
		fmt.Print("Content cached\n")
		response := Response{}
		err := json.Unmarshal(content, &response)
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
	fmt.Print("Content NON cached\n")
	res, err := http.Get(url)
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
