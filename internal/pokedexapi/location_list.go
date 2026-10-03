package pokedexapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationPage(pageURL *string) (*LocationPage, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	// fetch the list of locations from the PokeAPI
	res, err := c.httpClient.Get(url)
	// return an error if the request fails
	if err != nil {
		return nil, fmt.Errorf("error fetching location list: %v", err)
	}

	defer res.Body.Close()

	// read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// check for non-200 status codes
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching location list: status code %d", res.StatusCode)
	}

	// stream the JSON response into a LocationPage struct
	var locationPage LocationPage
	err = json.Unmarshal(body, &locationPage)

	if err != nil {
		return nil, fmt.Errorf("error decoding location page JSON: %v", err)
	}

	return &locationPage, nil

}
