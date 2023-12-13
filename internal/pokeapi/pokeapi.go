package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationsResponse struct {
	Count   int
	Next    *string
	Previous    *string
	Results []Location
}

type Location struct {
	Name string
	Url  string
}

func (client *Client) ListNextLocations(url *string) (LocationsResponse, error) {
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return LocationsResponse{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return LocationsResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
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

func (client *Client) ListPrevLocations(url *string) (LocationsResponse, error) {
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return LocationsResponse{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return LocationsResponse{}, nil
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResponse{}, nil
	}

	locations := LocationsResponse{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationsResponse{}, nil
	}

	return locations, nil
}