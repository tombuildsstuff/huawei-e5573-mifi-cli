package mifi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type WifiSettings struct {
	SSID    string
	Country string
}

type basicSettingsResponse struct {
	WifiCountry string `xml="WifiCountry"`
	WifiSsid    string `xml="WifiSsid"`
}

func (m Mifi) WifiSettings() (*WifiSettings, error) {
	endpoint := fmt.Sprintf("%s/api/wlan/basic-settings", m.Endpoint)
	resp, err := m.makeGetRequest(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving BasicSettings from Mifi: %s", err)
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bsr := &basicSettingsResponse{}
	err = xml.Unmarshal(responseData, bsr)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling XML: %+v", err)
	}

	if bsr != nil {
		result := WifiSettings{
			Country: bsr.WifiCountry,
			SSID:    bsr.WifiSsid,
		}
		return &result, nil
	}

	return nil, fmt.Errorf("XML wasn't valid: %s", string(responseData))
}
