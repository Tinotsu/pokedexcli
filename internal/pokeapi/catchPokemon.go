//Package pokeapi
package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"github.com/Tinotsu/pokedexcli/internal/pokecache"
)

func CatchPokemon (pokemon string, cache *pokecache.Cache, pokedex *map[string]PokemonDetails) string {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	content, isCached := cache.Get(url)
	if isCached {
	fmt.Print("CatchPokemon cached\n")
		pokemonDetails := PokemonDetails{}
		err := json.Unmarshal(content, &pokemonDetails)
		if err != nil {
			fmt.Print("CatchPokemon error:")
			log.Fatal(err)
		}
		return tryCatch(pokemonDetails, pokedex)
	}
	fmt.Print("CatchPokemon NON cached\n")
	res, err := http.Get(url)
	if err != nil {
		fmt.Print("CatchPokemon error: res error\n")
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Print("CatchPokemon error: wrong url ?\n")
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Print("CatchPokemon error: body error\n")
		log.Fatal(err)
	}
	cache.Add(url, body)

	pokemonDetails := PokemonDetails{}
	err = json.Unmarshal(body, &pokemonDetails)
	if err != nil {
		fmt.Print("CatchPokemon error: json.Unmarshal fail\n")
		log.Fatal(err)
	} 
	return tryCatch(pokemonDetails, pokedex)
}


func tryCatch(pokemon PokemonDetails, pokedex *map[string]PokemonDetails) string {
	baseExp := pokemon.BaseExperience
	luck := rand.Intn(baseExp)
	var result string
	if luck > 50 {
		result = fmt.Sprintf("%s was caught!", pokemon.Name)
		var inPokedex bool
		for k := range *pokedex {
			if k == pokemon.Name {
				inPokedex = true
			}
		}
		if !inPokedex {
			fmt.Println(pokemon.Name)
			poke := *pokedex
			poke[pokemon.Name] = pokemon
		}
		return result
	}
	result = fmt.Sprintf("%s escaped!", pokemon.Name)
	return result
}
