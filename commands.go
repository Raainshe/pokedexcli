package main

import (
	"fmt"
	"net/http"
	"os"
)

type mapsList struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []mapName `json:"results"`
}

type mapName struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap() error {

	res, err := http.Get()
	return nil
}

func commandMapB() error {

}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range CmdList {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
