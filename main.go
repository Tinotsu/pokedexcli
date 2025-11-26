package main

import (
	"fmt"
	"bufio"
	"os"
)

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
			case "map":
				showMap()
			case "mapb":
				showMapB()
			default:
				fmt.Print(str)
			}
			fmt.Print("\nPokedex > ")
		}		
	}
}
