package main

import (
	"errors"
	"fmt"
	"os"
)

// these functions used by keywords
func commandHelp(cfg *config, args ...string) error {
	fmt.Println("to the Pokedex!")
	fmt.Println("Usage:")

	for _, word := range isKeyword() {
		fmt.Println(word)
		fmt.Println(isKeyword()[word.description])
	}
	return nil
}

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, args ...string) error {
	response, err := cfg.pokeapiClient.GetMap(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	fmt.Println("Local areas:")
	for _, area := range response.Results {
		fmt.Printf(" _ %s\n", area.Name)
	}

	cfg.nextLocationURL = response.Next
	cfg.prevLocationURL = response.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	response, err := cfg.pokeapiClient.GetMap(cfg.prevLocationURL)
	if err != nil {
		return err
	}
	fmt.Println("Local areas:")
	for _, area := range response.Results {
		fmt.Printf(" _ %s \n", area.Name)
	}

	cfg.nextLocationURL = response.Next
	cfg.prevLocationURL = response.Previous
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.LocationExplore(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
