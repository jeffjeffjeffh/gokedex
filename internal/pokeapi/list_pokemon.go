package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)



func (client *Client) ListPokemon(area string) ([]string, error) {
	const BASE_URL = "https://pokeapi.co/api/v2/location-area/"
	url := BASE_URL + area

	if data, ok := client.cache.Get(&url); ok {
		res := PokemonListResponse{}
		
		err := json.Unmarshal(data, &res)
		if err != nil {
			return []string{}, err
		}

		nameList := []string{}
		
		for _, pokemon := range res.PokemonEncounters {
			nameList = append(nameList, pokemon.Pokemon.Name)
		}

		return nameList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}

	client.cache.Add(url, data)

	pokemonResponse := PokemonListResponse{}
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return []string{}, err
	}

	nameList := []string{}
		
	for _, pokemon := range pokemonResponse.PokemonEncounters {
		nameList = append(nameList, pokemon.Pokemon.Name)
	}

	return nameList, nil
}