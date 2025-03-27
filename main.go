package main

import (
	"time"

	"github.com/tomanta/pokedexcli/internal/pokeAPI"
)

func main() {
	pokeClient := pokeAPI.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokedex:       make(map[string]pokeAPI.PokemonDetails),
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
