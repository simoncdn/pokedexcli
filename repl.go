package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := scanner.Text()
		userCommand := cleanInput(userInput)[0]

		fmt.Printf("Your command was: %s\n", userCommand)
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}
