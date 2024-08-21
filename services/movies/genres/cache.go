package genres

import (
	"sync"

	"bookmymovie.app/bookmymovie/database"
)

type cache struct {
	genres []database.MoviesGenre
	mu      sync.RWMutex
}

func (c *cache) refresh(genres []database.MoviesGenre) {
	c.mu.Lock()
	c.genres = genres
	c.mu.Unlock()
}
