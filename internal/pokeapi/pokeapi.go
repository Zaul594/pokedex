package pokeapi

import (
	"net/http"
	"time"

	"github.com/Zaul594/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, casheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(casheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
