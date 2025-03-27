package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		// NOTE: BootDev code returns an error here wth a similar message
		fmt.Println("Which pokemon do you want to catch?")
		return nil
	}

	pokemonDetailsResp, err := cfg.pokeapiClient.PokemonDetails(args[0])
	if err != nil {
		// TODO: If lookup fails, the error appears here
		// "invalid character 'N' looking for beginning of value"
		return err
	}

	pokemon := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	chance := 0.9 - (float64(pokemonDetailsResp.BaseExperience) / 1000)
	if rand.Float64() > chance {
		fmt.Printf("%s escaped!\n", pokemon)
	} else {
		fmt.Printf("%s was caught!\n", pokemon)

		_, exists := cfg.pokedex[pokemon]
		if !exists {
			cfg.pokedex[pokemon] = pokemonDetailsResp
		}

	}
	return nil
}
