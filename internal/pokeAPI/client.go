package pokeAPI

import (
	"net/http"
	"time"

	"github.com/tomanta/pokedexcli/internal/pokecache"
)

type Client struct {
	cache 	   pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}