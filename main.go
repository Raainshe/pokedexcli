package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var CmdList map[string]cliCommand

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		input := cleanInput(text)
		if len(input) > 0 {
			if CmdList[input[0]].callback() != nil {
				fmt.Println("Invalid command")
			}
		}
	}

}

func init() {
	CmdList = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name: "map",
			description: "Explore the map of the pokemon world",
			callback: commandMap,
		}
	}
}

func cleanInput(text string) []string {
	var result []string

	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	result = strings.Fields(text)

	return result
}
