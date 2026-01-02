package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonStruct, error) {
	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	url := baseURL + pokemonName
	if data, ok := c.cache.Get(url); ok {
		fmt.Println("Found in cache!")
		pokemonStruct := PokemonStruct{}
		err := json.Unmarshal(data, &pokemonStruct)
		if err != nil {
			return PokemonStruct{}, err
		}
		return pokemonStruct, nil
	}
	fmt.Printf("Not found in Cache - fetching data\n")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonStruct{}, err
	}
	response, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonStruct{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusNotFound {
		return PokemonStruct{}, fmt.Errorf("Pokemon '%s' not found\n", pokemonName)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return PokemonStruct{}, err
	}
	c.cache.Add(url, data)
	pokemonStruct := PokemonStruct{}
	err = json.Unmarshal(data, &pokemonStruct)
	if err != nil {
		return PokemonStruct{}, err
	}
	return pokemonStruct, nil
}
