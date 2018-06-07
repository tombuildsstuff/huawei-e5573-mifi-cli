package mifi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type TrafficStatistics struct {
	SecondsConnectedToNetwork int
	DownloadedMB              float32
	UploadedMB                float32
}

type trafficStatisticsResponse struct {
	CurrentConnectTime int   `xml="CurrentConnectTime"`
	CurrentUpload      int32 `xml="CurrentUpload"`
	CurrentDownload    int32 `xml="CurrentDownload"`
}

func (m Mifi) TrafficStatistics() (*TrafficStatistics, error) {
	endpoint := fmt.Sprintf("%s/api/monitoring/traffic-statistics", m.Endpoint)
	resp, err := m.makeGetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tsr := &trafficStatisticsResponse{}
	err = xml.Unmarshal(responseData, tsr)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling XML: %+v", err)
	}

	if tsr != nil {
		mbDownloaded := bytesToMB(tsr.CurrentDownload)
		mbUploaded := bytesToMB(tsr.CurrentUpload)

		result := TrafficStatistics{
			SecondsConnectedToNetwork: tsr.CurrentConnectTime,
			DownloadedMB:              mbDownloaded,
			UploadedMB:                mbUploaded,
		}
		return &result, nil
	}

	return nil, fmt.Errorf("XML wasn't valid: %s", string(responseData))
}

func bytesToMB(input int32) float32 {
	return float32(input) / 1024 / 1024
}
