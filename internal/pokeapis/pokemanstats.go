package pokeapis

import (
	"encoding/json"
	"fmt"
)

func updateStats() {

}

func ListPokemonStats(pokemonName string) ([]byte, error) {
	// Verify whether pokeName is present in the map or not
	stats, ok := c.Get(pokemonName)
	if ok {
		exp := Experience{}
		err := json.Unmarshal(stats, &exp)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Name: %s\n", exp.Name)
		fmt.Printf("Height: %v\n", exp.Height)
		fmt.Printf("Weight: %v\n", exp.Weight)
		fmt.Println("Stats:")
		for _, val := range exp.Stats {
			fmt.Printf("  -%s: %v\n", val.SecondStat.StatName, val.BaseStat)
		}
		fmt.Println("Types:")
		for _, val := range exp.Types {
			fmt.Printf("  - %s\n", val.PType.Name)
		}
		return stats, nil
	}
	expectedMsg := fmt.Sprintf("you have not caught that pokemon")
	fmt.Println(expectedMsg)
	return []byte(expectedMsg), nil
}
