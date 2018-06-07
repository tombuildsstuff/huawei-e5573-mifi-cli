package mifi

import (
	"fmt"
	"net/http"
)

func (m *Mifi) ParseCookie() error {
	endpoint := fmt.Sprintf("%s/html/index.html", m.Endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		return fmt.Errorf("Error loading Cookie: %+v", err)
	}

	if cookies := resp.Cookies(); cookies != nil {
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "SessionID" {
				m.authorizationCookie = cookie
				return nil
			}
		}
	}

	return fmt.Errorf("Unable to find SessionID Cookie in response!")
}
