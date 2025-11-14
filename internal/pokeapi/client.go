package pokeapi

import (
	"net/http"
	"time"

	"github.com/eniolaomotee/Pokedex-Go/internal/pokecache"
)

// Client represents a PokeAPI client with caching capabilities.
type Client struct{
	cache *pokecache.Cache
	httpClient http.Client
}

// NewClient creates a new PokeAPI client with the specified cache duration.
func NewClient(timeout, cacheInterval time.Duration) *Client{
	return &Client{
		cache : pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}

}