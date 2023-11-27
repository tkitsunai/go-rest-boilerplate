package driver

import "net/http"

// HttpClient is embedded default http.client
type HttpClient struct {
	http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}
