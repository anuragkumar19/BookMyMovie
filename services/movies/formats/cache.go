package formats

import (
	"sync"

	"bookmymovie.app/bookmymovie/database"
)

type cache struct {
	formats []database.MoviesFormat
	mu      sync.RWMutex
}

func (c *cache) refresh(formats []database.MoviesFormat) {
	c.mu.Lock()
	c.formats = formats
	c.mu.Unlock()
}
