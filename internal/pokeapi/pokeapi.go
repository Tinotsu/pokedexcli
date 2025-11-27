// Package pokeapi
package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	"github.com/Tinotsu/pokedexcli/internal/pokecache"
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
type RegionDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
//-------------------------------------
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
func ExploreArea (location string, cache *pokecache.Cache) []string {
	fmt.Printf("Exploring %s...\n", location)
	url := "https://pokeapi.co/api/v2/location-area/" + location
	content, isCached := cache.Get(url)
	if isCached {
	// 	fmt.Print("exploreArea cached\n")
		regionDetails := RegionDetails{}
		err := json.Unmarshal(content, &regionDetails)
		if err != nil {
			fmt.Print("exploreArea error:")
			log.Fatal(err)
		}
		pokemons := []string{}
		for i := 0; i < len(regionDetails.PokemonEncounters); i++ {
			pokemons = append(pokemons, regionDetails.PokemonEncounters[i].Pokemon.Name)
		}
		return pokemons
	}
	// fmt.Print("exploreArea NON cached\n")
	res, err := http.Get(url)
	if err != nil {
		fmt.Print("exploreArea error: res error\n")
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Print("exploreArea error:")
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Print("exploreArea error: body error\n")
		log.Fatal(err)
	}
	cache.Add(url, body)

	regionDetails := RegionDetails{}
	err = json.Unmarshal(body, &regionDetails)
	if err != nil {
		fmt.Print("exploreArea error:")
		log.Fatal(err)
	}
	pokemons := []string{}
	for i := 0; i < len(regionDetails.PokemonEncounters); i++ {
		pokemons = append(pokemons, regionDetails.PokemonEncounters[i].Pokemon.Name)
	}

	return pokemons
}
func GetNextAreas(offset int, cache *pokecache.Cache) []string {
	strOffset := fmt.Sprintf("%d",offset)
	url := "https://pokeapi.co/api/v2/location-area/?offset=" + strOffset
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
	cache.Add(url, body)

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
