package mifi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Status struct {
	CurrentBatteryPercentage     int
	MaxSignalBars                int
	CurrentSignalBars            int
	NumberOfUsersConnectedToWifi int
}

type statusResponse struct {
	BatteryPercent  int `xml="BatteryPercent"`
	CurrentWifiUser int `xml="CurrentWifiUser"`
	SignalIcon      int `xml="SignalIcon"`
}

func (m Mifi) CurrentStatus() (*Status, error) {
	endpoint := fmt.Sprintf("%s/api/monitoring/status", m.Endpoint)
	resp, err := m.makeGetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sr := &statusResponse{}
	err = xml.Unmarshal(responseData, sr)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling XML: %+v", err)
	}

	if sr != nil {
		result := Status{
			CurrentBatteryPercentage:     sr.BatteryPercent,
			CurrentSignalBars:            sr.SignalIcon,
			MaxSignalBars:                5,
			NumberOfUsersConnectedToWifi: sr.CurrentWifiUser,
		}
		return &result, nil
	}

	return nil, fmt.Errorf("XML wasn't valid: %s", string(responseData))
}
