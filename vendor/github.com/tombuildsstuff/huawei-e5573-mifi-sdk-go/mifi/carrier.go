package mifi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Carrier struct {
	FullName  string
	ShortName string
	CarrierID int
}

type currentPLMNResponse struct {
	FullName  string `xml="FullName"`
	ShortName string `xml="ShortName"`
	Numeric   int    `xml="Numeric"`
}

func (m Mifi) CarrierDetails() (*Carrier, error) {
	endpoint := fmt.Sprintf("%s/api/net/current-plmn", m.Endpoint)
	resp, err := m.makeGetRequest(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Error making request for Current PLM: %s", err)
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	plmnResp := &currentPLMNResponse{}
	err = xml.Unmarshal(responseData, plmnResp)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling XML: %+v", err)
	}

	if plmnResp != nil {
		result := Carrier{
			CarrierID: plmnResp.Numeric,
			FullName:  plmnResp.FullName,
			ShortName: plmnResp.ShortName,
		}
		return &result, nil
	}

	return nil, fmt.Errorf("XML wasn't valid: %s", string(responseData))
}
