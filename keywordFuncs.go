package main

import (
	"errors"
	"fmt"
	"math/rand"
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

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a pokemon")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", fmt.Sprint(pokemon.Name))
	chance := rand.Intn(3)
	if chance == 1 {
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a pokemon")
	}

	for _, j := range cfg.pokedex {
		if j.Name == args[0] {
			fmt.Printf("Name:%s", j.Name)
			fmt.Printf("Height:%d", j.Height)
			fmt.Printf("Weight:%d", j.Weight)
			fmt.Println("Stats:")
			for _, i := range j.Stats {
				fmt.Printf("	-%s:%d \n", i.Stat.Name, i.BaseStat)
			}
			fmt.Println("Types:")
			for _, i := range j.Types {
				fmt.Printf("	-%s \n", i.Type.Name)
			}
			return nil
		}
	}
	fmt.Println("you do not have this pokemon")
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, i := range cfg.pokedex {
		fmt.Printf("	-%s \n", i.Name)
	}
	return nil
}
