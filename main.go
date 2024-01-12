package main

import (
	"time"

	"github.com/peterjunpark/gokedex/pokeapi"
)

type config struct {
	nextLocationsUrl *string
	prevLocationsUrl *string
	caughtPokemon    map[string]pokeapi.Pokemon
	pokeapiClient    pokeapi.Client
}

func main() {
	c := config{
		pokeapiClient: pokeapi.InitClient(time.Minute * 5),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&c)
}
