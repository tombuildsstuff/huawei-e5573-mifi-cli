package mifi

import (
	"fmt"
	"net/http"
)

func (m *Mifi) handleAuthorization() error {
	endpoint := fmt.Sprintf("%s/html/index.html", m.Endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	// if we've already got a Session ID, reuse it
	if m.authorizationCookie != nil {
		req.AddCookie(m.authorizationCookie)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error loading Cookie: %+v", err)
	}

	if m.authorizationCookie == nil {
		err = m.findAuthToken(resp.Cookies())
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Mifi) findAuthToken(input []*http.Cookie) error {
	if cookies := input; cookies != nil {
		for _, cookie := range input {
			if cookie.Name == "SessionID" {
				m.authorizationCookie = cookie
				return nil
			}
		}
	}

	return fmt.Errorf("Unable to find SessionID Cookie in response!")
}