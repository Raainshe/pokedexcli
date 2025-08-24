package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	url := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

	if Maps.Next != "" {
		url = Maps.Next
	}
	//check if what we want to call exists in Cache
	if data, exists := APICache.Get(url); exists {
		//use cache data instead
		err := json.Unmarshal(data, &Maps)
		if err != nil {
			return fmt.Errorf("failed to parse Json data from cache")
		}
		err = printMaps()
		if err != nil {
			return fmt.Errorf("error printing map: %w", err)
		}
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("status Code error: %d", res.StatusCode)
	}
	//add data to cache
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not conver cache data to []byte: %w", err)
	}
	APICache.Add(url, data)
	err = json.Unmarshal(data, &Maps)
	if err != nil {
		return fmt.Errorf("error decoding response data: %w", err)
	}

	err = printMaps()
	if err != nil {
		return fmt.Errorf("could not print maps: %w", err)
	}
	return nil
}

func commandMapB() error {

	if Maps.Previous == "" {
		return fmt.Errorf("there is no previous page")
	}
	//check if what we want to call exists in Cache
	if data, exists := APICache.Get(Maps.Previous); exists {
		//use cache data instead
		err := json.Unmarshal(data, &Maps)
		if err != nil {
			return fmt.Errorf("failed to parse Json data from cache")
		}
		err = printMaps()
		if err != nil {
			return fmt.Errorf("error printing map: %w", err)
		}
		return nil
	}
	url := Maps.Previous
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("status Code error: %d", res.StatusCode)
	}
	//add data to cache
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not conver cache data to []byte: %w", err)
	}
	APICache.Add(url, data)
	err = json.Unmarshal(data, &Maps)
	if err != nil {
		return fmt.Errorf("error decoding response data: %w", err)
	}

	err = printMaps()
	if err != nil {
		return fmt.Errorf("could not print maps: %w", err)
	}
	return nil

}

func printMaps() error {
	if Maps.Results == nil {
		return fmt.Errorf("nothing to print")
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
