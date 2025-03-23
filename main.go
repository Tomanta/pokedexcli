package main

import (
	"time"

	"github.com/tomanta/pokedexcli/internal/pokeAPI"
)

func main() {
	pokeClient := pokeAPI.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
