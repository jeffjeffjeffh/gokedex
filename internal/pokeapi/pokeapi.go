package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BASE_URL string = "https://pokeapi.co/api/v2/location-area?limit=20&offset="

type LocationsResponse struct {
	Count   int
	Next    *string
	Prev    *string
	Results []Location
}

type Location struct {
	Name string
	Url  string
}

func (client *Client) ListLocations(nextUrl *string) (LocationsResponse, error) {
	url := BASE_URL + "0"
	if nextUrl != nil {
		url = *nextUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsResponse{}, err
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return LocationsResponse{}, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("io readall error")
		return LocationsResponse{}, err
	}

	locations := LocationsResponse{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println("json unmarshall error")
		return LocationsResponse{}, err
	}

	return locations, nil
}