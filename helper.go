package main

import(
	"github.com/Tinotsu/pokedexcli/internal/pokeapi"
	"fmt"
)
func inspect (pokemon string, pokedex *map[string]pokeapi.PokemonDetails) {
	poke := *pokedex
	var inPokedex bool
	for k := range poke {
		if pokemon == k {
			inPokedex = true
		} 
	}
	if inPokedex {
		fmt.Printf("Name: %s\n", poke[pokemon].Name)
		fmt.Printf("Height: %d\n", poke[pokemon].Height)
		fmt.Printf("Weight: %d\n", poke[pokemon].Weight)
		fmt.Printf("Types: \n")
		for _, t := range poke[pokemon].Types {
			fmt.Printf("  - %s", t.Type.Name)
		}
	} else {
		fmt.Print("you have not caught that pokemon")
	}
}
