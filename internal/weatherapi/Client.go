package weatherapi

import (
	"net/http"
	"time"

	"github.com/arturogood17/forecast/internal/weathercache"
)

type Client struct {
	Cache      weathercache.Cache
	httpClient *http.Client
}

func NewClient(timeout, interval time.Duration) Client {
	c := Client{
		Cache: weathercache.WCache(interval),
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
	return c
}
