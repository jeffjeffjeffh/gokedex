package pokeapi

import (
	"net/http"
	"time"
)

type Client struct{
	httpClient http.Client
	cache PokeCache
}

func NewClient(cxnTimeout, cacheInterval, cacheTimeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: cxnTimeout,
		},
		cache: NewCache(cacheInterval, cacheTimeout),
	}
}