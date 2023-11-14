package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// startRepl starts a ui that scanns the users input for key words then displays an ouput dependong on the key words inputed by the user.
func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		cfg.location = words[1]

		command, exists := isKeyword()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// cleanInput makes the words in the input all lowercase so the Pokedex can understand the key words no mater how they are inputed.
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
