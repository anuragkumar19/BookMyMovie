package languages

import (
	"sync"

	"bookmymovie.app/bookmymovie/database"
)

type cache struct {
	languages []database.MoviesLanguage
	index     map[int64]database.MoviesLanguage
	mu        sync.RWMutex
}

func (c *cache) refresh(langs []database.MoviesLanguage) {
	c.mu.Lock()
	m := make(map[int64]database.MoviesLanguage, len(langs))
	for _, l := range langs {
		m[l.ID] = l
	}
	c.languages = langs
	c.index = m

	c.mu.Unlock()
}
