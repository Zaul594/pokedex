package main

import (
	"fmt"
	"os"
)

// these functions used by keywords
func commandHelp() error {
	fmt.Println("to the Pokedex!")
	fmt.Println("Usage:")

	for _, word := range isKeyword() {
		fmt.Println(word)
		fmt.Println(isKeyword()[word.description])
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {

	for i := 1; i <= 20; i++ {
		locationID, location, err := getMap(i)
		if err != nil {
			return err
		}
		fmt.Println(locationID, location)
	}
	return nil
}

func commandMapb() error {

	for i := 20; i >= 1; i-- {
		locationID, location, err := getMap(i)
		if err != nil {
			return err
		}
		fmt.Println(locationID, location)
	}
	return nil
}
