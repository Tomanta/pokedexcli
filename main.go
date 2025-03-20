package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	result := []string{}
	split_string := strings.Fields(text)
	for i := range split_string {
		result = append(result, strings.ToLower(split_string[i]))
	}

	return result
}

