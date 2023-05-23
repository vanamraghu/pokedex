package pokeapis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	PokeName string `json:"name"`
	PokeUrl  string `json:"url"`
}

type PokemonEncounters struct {
	Poke Pokemon `json:"pokemon"`
}

type PokemonLocation struct {
	Encounters []PokemonEncounters `json:"pokemon_encounters"`
	Name       string              `json:"name"`
}

var pokeManData *PokemonLocation

func ListPokemonLocations(locationArea string) (PokemonLocation, error) {
	url := baseURL + "/location-area" + "/" + locationArea
	// Verify whether url is present in the cache
	val, ok := c.Get(url)
	// If present, read from cache
	if ok {
		location := PokemonLocation{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return PokemonLocation{}, err
		}
		return location, nil
	}

	// Use http get to get the response
	response, err := http.Get(url)
	if err != nil {
		return PokemonLocation{}, err
	}
	data, err := io.ReadAll(response.Body)
	err = response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &pokeManData)
	if err != nil {
		fmt.Println(err)
	}
	c.Add(url, data)
	return *pokeManData, nil
}
