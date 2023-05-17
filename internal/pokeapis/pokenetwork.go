package pokeapis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokecache"
	"sync"
)

const LOCATION = "https://pokeapi.co/api/v2/location-area"

type LocationStructure struct {
	LocationCount int               `json:"count"`
	NextUrl       *string           `json:"next"`
	PreviousUrl   *string           `json:"previous"`
	Results       []locationDetails `json:"results"`
}

type locationDetails struct {
	LocationName string `json:"name"`
	Url          string `json:"url"`
}

type Config struct {
	nextLocationURL *string
	prevLocationURL *string
}

var responseData *LocationStructure

//var c = pokecache.NewCache(5 * time.Second)

var c = pokecache.Cache{
	Mux:       &sync.Mutex{},
	CacheData: make(map[string]pokecache.CacheEntry),
}
var baseURL = "https://pokeapi.co/api/v2"

func GetLocationData(pageUrl *string) (LocationStructure, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	val, ok := c.Get(url)
	if ok {
		locationData := LocationStructure{}
		err := json.Unmarshal(val, &locationData)
		if err != nil {
			return LocationStructure{}, err
		}
		return locationData, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationStructure{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return LocationStructure{}, err
	}
	err = res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return LocationStructure{}, err
	}
	c.Add(url, body)
	return *responseData, nil
}

func addCachedData(url string, locationData LocationStructure, data []byte) []byte {
	for _, val := range locationData.Results {
		data = append(data, []byte(val.LocationName+"\n")...)
	}
	c.Add(url, data)
	return data
}

func displayCachedData(data []byte) {
	fmt.Printf("%s\n", data)
}
