package languages

import (
	"sync"

	"bookmymovie.app/bookmymovie/database"
)

type cache struct {
	languages []database.MoviesLanguage
	mu        sync.RWMutex
}

func (c *cache) refresh(langs []database.MoviesLanguage) {
	c.mu.Lock()
	c.languages = langs
	c.mu.Unlock()
}
