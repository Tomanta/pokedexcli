package main

import (
	"fmt"
	"github.com/tomanta/pokedexcli/internal/pokeAPI"
)

func commandMap(cfg *config) error {
	res, err := pokeAPI.GetLocationAreas(cfg.next)
	if err != nil {
		return err
	}

	for i := range res {
		fmt.Println(res[i])
	}

	return nil
}

func commandMapb(cfg *config) error {
	fmt.Println("TO BE IMPLEMENTED")

	return nil
}
