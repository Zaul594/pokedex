package main

import (
	"time"

	"github.com/Zaul594/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	location        string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(10*time.Second, time.Minute*5),
	}
	startRepl(&cfg)
}
