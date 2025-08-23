package main

import (
	"encoding/json"
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

var Maps mapsList

func commandMap() error {

	url := "https://pokeapi.co/api/v2/location-area/"

	if Maps.Next != "" {
		url = Maps.Next
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error fetching data: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Status Code error: %d", res.StatusCode)
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&Maps)
	if err != nil {
		return fmt.Errorf("Error decoding response data: %w", err)
	}
	err = printMaps()
	if err != nil {
		return fmt.Errorf("Could not print maps: %w", err)
	}
	return nil
}

func commandMapB() error {

	if Maps.Previous == "" {
		return fmt.Errorf("There is no previous page")
	}
	url := Maps.Previous
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error fetching data: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Status Code error: %d", res.StatusCode)
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&Maps)
	if err != nil {
		return fmt.Errorf("Error decoding response data: %w", err)
	}
	err = printMaps()
	if err != nil {
		return fmt.Errorf("Could not print maps: %w", err)
	}
	return nil

}

func printMaps() error {
	if Maps.Results == nil {
		return fmt.Errorf("Nothing to print")
	}
	for _, smap := range Maps.Results {
		fmt.Printf("%s\n", smap.Name)
	}
	return nil
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
