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