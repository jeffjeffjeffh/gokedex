package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) Catch(name string) (PokemonResponse, error) {
	const BASE_URL = "https://pokeapi.co/api/v2/pokemon/"
	url := BASE_URL + name
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	if res.StatusCode == 404 {
		return PokemonResponse{}, fmt.Errorf("a pokemon with the name %s was not found", name)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemonResponse := PokemonResponse{}
	
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return PokemonResponse{}, err
	}

	return pokemonResponse, nil
}