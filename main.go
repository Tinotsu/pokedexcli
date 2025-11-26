package main

import (
	"fmt"
	"bufio"
	"os"
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

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		for scanner.Scan(){
			msg := scanner.Text()
			str := cleanInput(msg)
			switch msg {
			case "exit" :
				commandExit()
			case "help":
				showHelp()
			default:
				fmt.Print(str)
			}
			fmt.Print("\nPokedex > ")
		}		
	}
}
