package pokeapi

import (
	"net/http"
	"time"

	"github.com/davidheeren/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache   *pokecache.Cache
}

func NewClient(httpTimeout, cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: httpTimeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
