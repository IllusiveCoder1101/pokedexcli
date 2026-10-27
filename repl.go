package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))

}

func repl(a *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		userInput := ""
		for scanner.Scan() {
			userInput = scanner.Text()
			break
		}
		userCommand := cleanInput(userInput)
		for _, command := range a.commands {
			userCommandName := userCommand[0]

			if userCommandName == command.name {
				if command.name == "explore" || command.name == "catch" || command.name == "inspect" {
					userCommandFlag := userCommand[1]
					command.callback(a, userCommandFlag)
				} else {
					command.callback(a, "")
				}

				break
			}
		}

	}
}
