package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for i := range cfg.pokedex {
		fmt.Printf("  - %s\n", cfg.pokedex[i].Name)
	}
	
	return nil
}