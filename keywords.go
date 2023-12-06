package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

// sets up key words and what they do when entered into the pokedex.
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
		"explore": {
			name:        "explore <location_name>",
			description: "shows the pokemon that can be found in this area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "command used to try to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "command used to see the information of a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "command used to show pokemon caught by the player",
			callback:    commandPokedex,
		},
	}
	return keywords
}
