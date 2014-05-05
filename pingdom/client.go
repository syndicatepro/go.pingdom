// Steve Phillips / elimisteve
// 2014.05.05

package pingdom

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	appKey   string
	email    string
	password string
}

func NewClient(appKey, email, password string) *Client {
	return &Client{appKey, email, password}
}

func (c *Client) get(url string) ([]byte, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("App-Key", c.appKey)
	req.SetBasicAuth(c.email, c.password)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error GETing %s: %v", url, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading resp.Body: %v", err)
	}
	defer resp.Body.Close()

	return body, nil
}

func (c *Client) Checks() ([]*Check, error) {
	body, err := c.get(urlChecks)
	if err != nil {
		return nil, err
	}

	checks := &checksResponse{}
	err = json.Unmarshal(body, checks)
	if err != nil {
		return nil, fmt.Errorf("Error Unmarshaling Checks: %v", err)
	}

	if checks.Error != nil {
		return nil, fmt.Errorf("%v %v: %v", checks.Error.StatusCode,
			checks.Error.StatusDesc, checks.Error.ErrorMessage)
	}

	return checks.Checks, nil
}
