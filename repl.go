package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 

		words := cleanInput(scanner.Text())
		
		if len(words) == 0 {
			continue
		}
		fmt.Println("Your command was:", words[0])
	}

}

func cleanInput(text string) []string {
	result := []string{}
	split_string := strings.Fields(text)
	for i := range split_string {
		result = append(result, strings.ToLower(split_string[i]))
	}

	return result
}

