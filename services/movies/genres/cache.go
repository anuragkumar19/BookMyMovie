package genres

import (
	"sync"

	"bookmymovie.app/bookmymovie/database"
)

type cache struct {
	genres []database.MoviesGenre
	index  map[int64]database.MoviesGenre
	mu     sync.RWMutex
}

func (c *cache) refresh(genres []database.MoviesGenre) {
	c.mu.Lock()
	m := make(map[int64]database.MoviesGenre, len(genres))
	for _, g := range genres {
		m[g.ID] = g
	}
	c.genres = genres
	c.index = m
	c.mu.Unlock()
}
