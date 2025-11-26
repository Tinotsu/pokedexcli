package main

import(
	"os"
	"fmt"
	"github.com/Tinotsu/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var listCommands map[string]cliCommand 
func init() {
	listCommands = map[string]cliCommand {
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:		"help",
			description: "Displays a help message",
			callback: 	showHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback: showMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback: showMapB,
		},
	}}

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
func showMap() error {
	locationsName := pokeapi.GetNextAreas(showMapOffset)
	for _, loc := range locationsName {
		fmt.Println(loc)
	}
	showMapOffset += 20
	return nil
}

func showMapB() error {
	showMapOffset -= 40
	locationsName := pokeapi.GetNextAreas(showMapOffset)
	for _, loc := range locationsName {
		fmt.Println(loc)
	}
	return nil
}
