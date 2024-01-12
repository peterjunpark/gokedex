package pokeapi

import (
	"net/http"
	"time"

	"github.com/peterjunpark/gokedex/datacache"
)

type Client struct {
	cache      datacache.Cache
	httpClient http.Client
}

func InitClient(cacheInterval time.Duration) Client {
	return Client{
		cache: datacache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
