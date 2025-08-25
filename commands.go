package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
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

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Weight         int    `json:"weight"`
	Stats          PokemonStats
	Types          []string
	RawStats       []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	RawTypes []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type PokemonStats struct {
	HP             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}

type LocationAreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

var Maps mapsList
var locationDetail LocationAreaDetail
var Pokedex map[string]Pokemon

func commandPokedex(_ []string) error {
	if len(Pokedex) == 0 {
		return fmt.Errorf("you have no pokemon. catch pokemond using 'catch <pokemon>'")
	}
	for name, _ := range Pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}

func commandInspect(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("inspect requires a pokemon name")
	}

	if poke, exists := Pokedex[args[0]]; exists {
		fmt.Printf("Name: %s\n", poke.Name)
		fmt.Printf("Height: %d\n", poke.Height)
		fmt.Printf("Weight: %d\n", poke.Weight)
		fmt.Println("Stats:")
		fmt.Printf("  -hp: %d\n", poke.Stats.HP)
		fmt.Printf("  -attack: %d\n", poke.Stats.Attack)
		fmt.Printf("  -defense: %d\n", poke.Stats.Defense)
		fmt.Printf("  -special-attack: %d\n", poke.Stats.SpecialAttack)
		fmt.Printf("  -special-defense: %d\n", poke.Stats.SpecialDefense)
		fmt.Printf("  -speed: %d\n", poke.Stats.Speed)
		fmt.Println("Types:")
		for _, t := range poke.Types {
			fmt.Printf("  - %s\n", t)
		}
		return nil
	} else {
		return fmt.Errorf("you have not caught that pokemon")
	}
}

func (p *Pokemon) populateStats() {
	// Convert stats to clean format
	for _, stat := range p.RawStats {
		switch stat.Stat.Name {
		case "hp":
			p.Stats.HP = stat.BaseStat
		case "attack":
			p.Stats.Attack = stat.BaseStat
		case "defense":
			p.Stats.Defense = stat.BaseStat
		case "special-attack":
			p.Stats.SpecialAttack = stat.BaseStat
		case "special-defense":
			p.Stats.SpecialDefense = stat.BaseStat
		case "speed":
			p.Stats.Speed = stat.BaseStat
		}
	}

	// Convert types to clean format
	for _, t := range p.RawTypes {
		p.Types = append(p.Types, t.Type.Name)
	}
}

func (p *Pokemon) tryCatch() error {
	fmt.Printf("Throwing a Pokeball at %s...\n", p.Name)
	var catchChance int

	switch {
	case p.BaseExperience <= 80:
		catchChance = 80
	case p.BaseExperience <= 150:
		catchChance = 60
	case p.BaseExperience <= 250:
		catchChance = 40
	default:
		catchChance = 20
	}

	randomRoll := rand.Intn(100) + 1

	if randomRoll <= catchChance {
		fmt.Printf("%s was caught!\n", p.Name)
		Pokedex[p.Name] = *p
		return nil
	} else {
		fmt.Printf("%s escaped!\n", p.Name)
	}
	return nil
}

func commandCatch(args []string) error {
	var poke Pokemon
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]

	//check if we have it in cache before calling api
	data, exists := APICache.Get(url)
	if exists {
		err := json.Unmarshal(data, &poke)
		if err != nil {
			return fmt.Errorf("error unmarshalling cache data: %w", err)
		}
		//try to catch pokemon and add it
		poke.tryCatch()
	}
	//doesnt exist call api
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching pokemon: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("api error code: %d", res.StatusCode)
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading res.body for pokemon api call: %w", err)
	}
	//add it to cache
	APICache.Add(url, data)
	err = json.Unmarshal(data, &poke)
	if err != nil {
		return fmt.Errorf("error unmarshaling pokemon data: %w", err)
	}
	poke.tryCatch()
	return nil
}

func commandExplore(args []string) error {

	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	data, exists := APICache.Get(url)
	if exists {
		err := json.Unmarshal(data, &locationDetail)
		if err != nil {
			return fmt.Errorf("error unmarshaling json for location detail: %w", err)
		}

		err = printPokemonInLocation(args[0])
		if err != nil {
			return err
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("api error getting pokemon: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("status Code error: %d", res.StatusCode)
	}
	//add to data to cache
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading body %w", err)
	}
	APICache.Add(url, data)
	err = json.Unmarshal(data, &locationDetail)
	if err != nil {
		return fmt.Errorf("error unmarshaling json for location detail: %w", err)
	}

	err = printPokemonInLocation(args[0])
	if err != nil {
		return err
	}
	return nil
}

func commandMap(_ []string) error {
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

func commandMapB(_ []string) error {
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

func printPokemonInLocation(location string) error {
	fmt.Printf("Exploring %s...\n", location)
	if len(locationDetail.PokemonEncounters) == 0 {
		return fmt.Errorf("no pokemon found in this location")
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationDetail.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandHelp(_ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range CmdList {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit(_ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
