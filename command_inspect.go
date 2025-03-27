package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		// NOTE: BootDev code returns an error here wth a similar message
		fmt.Println("Which pokemon do you want to inspect?")
		return nil
	}

	name := args[0]

	data, exists := cfg.pokedex[name]
	if !exists {
		fmt.Printf("%s not yet caught!\n", name)
		return nil
	}

	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %d\n", data.Height)
	fmt.Printf("Weight: %d\n", data.Weight)

	stats := data.Stats
	fmt.Println("Stats:")
	for i := range stats {
		fmt.Printf("  -%s: %d\n", stats[i].Stat.Name, stats[i].BaseStat)
	}

	types := data.Types
	fmt.Println("Types")
	for i := range types {
		fmt.Printf("  - %s\n", types[i].Type.Name)
	}

	return nil
}
