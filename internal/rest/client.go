package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

type SegmentRegistrationRequest struct {
	// Empty request as node ID and name are assigned by the gateway
}

type SegmentRegistrationResponse struct {
	NodeID    int64  `json:"node_id"`    // Node ID assigned by the gateway
	NodeName  string `json:"node_name"`  // Temporary name assigned by the gateway
	SegmentID string `json:"segment_id"` // Segment where the node is placed
	Success   bool   `json:"success"`
}

type CertificateRequest struct {
	NodeID    int64  `json:"node_id"`
	NodeName  string `json:"node_name"`
	SegmentID string `json:"segment_id"`
}

type CertificateResponse struct {
	Certificate string `json:"certificate"`
	Success     bool   `json:"success"`
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// RegisterInSegment registers with the gateway to get assigned to a segment and receive a node ID and temporary name
func (c *Client) RegisterInSegment() (*SegmentRegistrationResponse, error) {
	reqBody := SegmentRegistrationRequest{}
	
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/segment/register", c.baseURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response SegmentRegistrationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response, nil
}

func (c *Client) GetQNECertificate(nodeID int64, nodeName, segmentID string) (*CertificateResponse, error) {
	reqBody := CertificateRequest{
		NodeID:    nodeID,
		NodeName:  nodeName,
		SegmentID: segmentID,
	}
	
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/certificate", c.baseURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response CertificateResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response, nil
}
