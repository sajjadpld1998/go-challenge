package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"user_segmentation_service/config"
)

type EsSegmentPain struct {
	UserId  string `json:"user_id"`
	Segment string `json:"segment"`
}

func (obj Http) MakeHTTPRequest(method string, url string, body []byte) (*http.Response, error) {
	// Create a new request with the specified method, URL, and body
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set appropriate headers if needed (e.g., JSON content type)
	req.Header.Set("Content-Type", "application/json")

	// Use the default HTTP client to make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}

	return resp, nil
}

func (obj Http) SendPairDataToES(userId, segment string) (err error) {
	body := EsSegmentPain{UserId: userId, Segment: segment}

	data, err := json.Marshal(body)
	if err != nil {
		return
	}

	resp, err := obj.MakeHTTPRequest("POST", config.GetConfig().Url.EstimateServiceStorePair, data)
	if err != nil {
		return
	}
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = fmt.Errorf("status code is %d", resp.StatusCode)
	}

	return
}
