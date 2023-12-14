package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) ListLocations(url *string) (LocationAreaResponse, error) {
	if data, ok := client.cache.Get(url); ok {
		res := LocationAreaResponse{}
		
		err := json.Unmarshal(data, &res)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		for _, loc := range res.Results {
		fmt.Println(loc.Name)
		}	
	}

	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("io readall error")
		return LocationAreaResponse{}, err
	}

	client.cache.Add(*url, data)

	locations := LocationAreaResponse{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println("json unmarshall error")
		return LocationAreaResponse{}, err
	}

	return locations, nil
}

func (client *Client) ListPokemon(area string) ([]string, error) {
	const BASE_URL = "https://pokeapi.co/api/v2/location-area/"
	url := BASE_URL + area

	if data, ok := client.cache.Get(&url); ok {
		res := PokemonResponse{}
		
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

	pokemonResponse := PokemonResponse{}
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