package infra

import (
	"net/http"
	"strings"
)

type NetworkClient interface {
	DoRequest(url string) (int, error)
}

type NetworkClientImpl struct{}

func NewNetworkClient() NetworkClient {
	return &NetworkClientImpl{}
}

func (n *NetworkClientImpl) DoRequest(url string) (int, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil && resp == nil {
		return -1, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
