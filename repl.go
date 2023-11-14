package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// startRepl starts a ui that scanns the users input for key words then displays an ouput dependong on the key words inputed by the user.
func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()
		text := cleanInput(scanner.Text())

		cfg.location = text[1]
		command := text[0]
		keyWord, exists := isKeyword()[command]
		if exists {
			err := keyWord.callback(cfg)
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
