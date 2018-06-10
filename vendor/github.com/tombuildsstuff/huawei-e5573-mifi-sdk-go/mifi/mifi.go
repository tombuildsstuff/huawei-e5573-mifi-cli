package mifi

import (
	"fmt"
	"net/http"
)

type Mifi struct {
	Endpoint string

	authorizationCookie *http.Cookie
}

func (m Mifi) makeGetRequest(endpoint string) (*http.Response, error) {
	if m.authorizationCookie == nil {
		err := m.handleAuthorization()
		if err != nil {
			return nil, fmt.Errorf("Error obtaining authentication cookie for Mifi: %+v", err)
		}
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(m.authorizationCookie)

	return http.DefaultClient.Do(req)
}