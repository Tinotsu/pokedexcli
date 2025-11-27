package main

import(
	"os"
	"fmt"
	"github.com/Tinotsu/pokedexcli/internal/pokeapi"
	"github.com/Tinotsu/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
type config struct {
    cache   *pokecache.Cache
    pokedex map[string]pokeapi.PokemonDetails
    // maybe: args []string
}

var listCommands map[string]cliCommand 
func init() {
	listCommands = map[string]cliCommand {
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    closure,
		},
		"help": {
			name:		"help",
			description: "Displays a help message",
			callback: 	closure,
		},
		"map": {
			name: "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback: closure,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback: closure,
		},
		"explore": {
			name: "explore",
			description: "Displays the different pokemons of the regions passed as argument",
			callback: closure,
		},
		"catch": {
			name: "catch",
			description: "Catch the pokemon passed as argument",
			callback: closure,
		},
		"pokedex": {
			name: "pokedex",
			description: "Display the pokemons in the user pokedex",
			callback: closure,
		},
	}}
func showPokedex (pokedex map[string]pokeapi.PokemonDetails) error {
    if len(pokedex) > 0 {
        for name := range pokedex {
            fmt.Println(name)
        }
    } else {
        fmt.Println("You caught 0 pokemon")
    }
    return nil
}
func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func showHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, value := range listCommands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

var showMapOffset int
func closure (cfg *config) error {return nil}
func showMap(cache *pokecache.Cache) error {
	locationsName := pokeapi.GetNextAreas(showMapOffset, cache)
	for _, loc := range locationsName {
		fmt.Println(loc)
	}
	showMapOffset += 20
	return nil
}

func showMapB(cache *pokecache.Cache) error {
	showMapOffset -= 40
	locationsName := pokeapi.GetNextAreas(showMapOffset, cache)
	for _, loc := range locationsName {
		fmt.Println(loc)
	}
	return nil
}
func explore(cache *pokecache.Cache, region string) error {
	pokemons := pokeapi.ExploreArea(region, cache)
	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
	return nil
}
func catch(cache *pokecache.Cache, pokemon string, pokedex map[string]pokeapi.PokemonDetails) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	result := pokeapi.CatchPokemon(pokemon, cache, &pokedex)
	fmt.Print(result)
	return nil
}
