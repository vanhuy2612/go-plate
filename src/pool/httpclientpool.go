package pool

import (
	"net/http"
	"time"
)

var HttpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:       100,
		MaxConnsPerHost:    100,
		IdleConnTimeout:    90 * time.Second,
		DisableCompression: false,
	},
	Timeout: 10 * time.Second,
}
