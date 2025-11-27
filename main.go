package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/Tinotsu/pokedexcli/internal/pokecache"
	"github.com/Tinotsu/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	interval := time.Second * 30
	cache := pokecache.NewCache(interval)
	pokedex := make(map[string]pokeapi.PokemonDetails)
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		for scanner.Scan(){
			msg := scanner.Text()
			str := cleanInput(msg)
			switch str[0] {
			case "explore":
				explore(cache, str[1])
			case "catch":
				catch(cache, str[1], pokedex)
			case "inspect":
				inspect(str[1], &pokedex)
			}
			switch msg {
			case "cache":
				pokecache.CacheTest()
			case "exit" :
				commandExit()
			case "help":
				showHelp()
			case "map":
				showMap(cache)
			case "mapb":
				showMapB(cache)
			case "pokedex":
				showPokedex(pokedex)
			}
			fmt.Print("\nPokedex > ")
		}		
	}
}
