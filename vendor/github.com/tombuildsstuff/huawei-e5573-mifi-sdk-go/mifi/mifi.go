package mifi

import (
	"fmt"
	"net/http"
)

type Mifi struct {
	Endpoint string
	SSID     string

	authorizationCookie *http.Cookie
}

func (m Mifi) makeGetRequest(endpoint string) (*http.Response, error) {
	if m.authorizationCookie == nil {
		return nil, fmt.Errorf("Authorization Cookie was not found!")
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(m.authorizationCookie)

	return http.DefaultClient.Do(req)
}
