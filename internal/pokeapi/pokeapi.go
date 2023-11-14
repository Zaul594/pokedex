package pokeapi

import (
	"net/http"
	"time"

	"github.com/Zaul594/pokedexcli/internal/pokecache"
)

// the struct for the Client
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewCLient is used to create a new client that rimes out after a sertain amount of time.
func NewClient(timeout, casheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(casheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
