package pokeapis

import (
	"encoding/json"
	"io"
	"net/http"
)

type Experience struct {
	BaseExperience int `json:"base_experience"`
}

var exp *Experience

func PokemonCatchDetails(pokemonName string) (Experience, error) {
	pokeUrl := baseURL + "/pokemon/" + pokemonName

	// Use http get to the response
	response, err := http.Get(pokeUrl)
	if err != nil {
		return Experience{}, err
	}
	// use io read to read the response body
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Experience{}, nil
	}
	err = response.Body.Close()
	if err != nil {
		return Experience{}, err
	}
	// Unmarshal the data based on the structure
	err = json.Unmarshal(data, &exp)
	if err != nil {
		return Experience{}, err
	}
	return *exp, nil

}
