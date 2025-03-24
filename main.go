package main

import (
	"time"

	"github.com/tomanta/pokedexcli/internal/pokeAPI"
)

func main() {
	pokeClient := pokeAPI.NewClient(5 * time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
