package pokeapis

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type Experience struct {
	BaseExperience int    `json:"base_experience"`
	Name           string `json:"name"`
}

type PokeMon struct {
	Experience int
	PokeName   string
}

var exp *Experience

const EXPERIENCE = 300

func PokemonCatchDetails(pokemonName string) (string, error) {
	status := ""
	pokeUrl := baseURL + "/pokemon/" + pokemonName
	fmt.Printf("Throwing a Pokeball at %s\n", pokemonName)
	// Find whether key is present
	val, ok := c.Get(pokemonName)
	if ok {
		return string(val), nil
	}

	// Use http get to the response
	response, err := http.Get(pokeUrl)
	if err != nil {
		return "", err
	}
	// use io read to read the response body
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", nil
	}
	err = response.Body.Close()
	if err != nil {
		return "", err
	}
	// Unmarshal the data based on the structure
	err = json.Unmarshal(data, &exp)
	if err != nil {
		return "", err
	}

	// Use math random
	randValue := rand.Intn(EXPERIENCE)
	actualExperience := exp.BaseExperience
	if randValue < actualExperience {
		status = fmt.Sprintf("%s was caught!\n", exp.Name)
		c.Add(exp.Name, []byte(status))
	} else {
		status = fmt.Sprintf("%s escaped!\n", exp.Name)
	}
	return status, nil
}
