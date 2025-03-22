package main

import (
	"fmt"
	"github.com/tomanta/pokedexcli/internal/pokeAPI"
)

func commandMap(cfg *config) error {
	res := pokeAPI.CallAPI()
	fmt.Println(res)

	return nil
}

func commandMapb(cfg *config) error {
	fmt.Println("TO BE IMPLEMENTED")

	return nil
}
