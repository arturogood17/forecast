package weatherapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) Climate(name ...string) (Current, error) {
	url := baseURL + "/current.json?key=" + apiKey + "&q="
	if len(name) == 0 {
		return Current{}, errors.New("please provide a location")
	}
	n := name[0]
	url += n
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Current{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Current{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Current{}, err
	}
	var current_climate Current
	if err := json.Unmarshal(data, &current_climate); err != nil {
		return Current{}, err
	}
	return current_climate, nil
}
