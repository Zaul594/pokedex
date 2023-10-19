package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func isKeyword() map[string]cliCommand {

	keywords := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "shows a list of the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "shows a list of the previos 20 locations",
			callback:    commandMapb,
		},
	}
	return keywords
}
