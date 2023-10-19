package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Location struct {
	Count int    `json:"count"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

func getMap(id int) (int, string, error) {
	pokeapi := "https://pokeapi.co/api/v2/location/" + strconv.Itoa(id) + "/"
	response, err := http.Get(pokeapi)
	if err != nil {
		return 0, "", err
	}
	locations, err := io.ReadAll(response.Body)
	location := Location{}
	err = json.Unmarshal(locations, &location)
	if err != nil {
		return 0, "", err
	}

	if err != nil {
		return 0, "", err
	}

	return location.ID, location.Name, nil
}
