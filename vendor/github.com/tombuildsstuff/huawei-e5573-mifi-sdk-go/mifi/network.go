package mifi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

var thirdGenNetworkMode = "0201"
var lteNetworkMode = "00"

type NetworkSettings struct {
	LTEEnabled bool
}

func (n NetworkSettings) NetworkMode() string {
	if n.LTEEnabled {
		return "4G/LTE Enabled"
	}

	return "4G/LTE Disabled"
}

type request struct {
	NetworkBand string `xml="NetworkBand"`
	NetworkMode string `xml="NetworkMode"`
	LTEBand     string `xml="LTEBand"`
}

func (m Mifi) NetworkSettings() (*NetworkSettings, error) {
	endpoint := fmt.Sprintf("%s/api/net/net-mode", m.Endpoint)
	resp, err := m.makeGetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	nr := &request{}
	err = xml.Unmarshal(responseData, nr)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling XML: %+v", err)
	}

	if nr != nil {
		result := NetworkSettings{
			LTEEnabled: nr.NetworkMode == lteNetworkMode,
		}
		return &result, nil
	}

	return nil, fmt.Errorf("XML wasn't valid: %s", string(responseData))
}
