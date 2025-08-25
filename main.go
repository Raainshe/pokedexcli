package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokecache "github.com/raainshe/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

var CmdList map[string]cliCommand
var APICache *pokecache.Cache

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		input := cleanInput(text)
		if len(input) > 0 {
			if cmd, exists := CmdList[input[0]]; exists {
				if err := cmd.callback(input[1:]); err != nil {
					fmt.Println("Command failed:", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}

}

func init() {
	APICache = pokecache.NewCache(10 * time.Minute)
	Pokedex = make(map[string]Pokemon)
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
			name:        "map",
			description: "Explore the map of the pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Explore the map of pokemod world {previous}",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "See Pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	var result []string

	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	result = strings.Fields(text)

	return result
}
