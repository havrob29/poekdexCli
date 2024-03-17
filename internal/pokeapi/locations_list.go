package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (AreaLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return AreaLocations{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return AreaLocations{}, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return AreaLocations{}, err
	}

	locations := AreaLocations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println(err)
		return AreaLocations{}, err
	}

	return locations, nil
}
