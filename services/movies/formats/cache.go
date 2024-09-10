package formats

import (
	"sync"

	"bookmymovie.app/bookmymovie/database"
)

type cache struct {
	formats []database.MoviesFormat
	index   map[int64]database.MoviesFormat
	mu      sync.RWMutex
}

func (c *cache) refresh(formats []database.MoviesFormat) {
	c.mu.Lock()
	m := make(map[int64]database.MoviesFormat, len(formats))
	for _, f := range formats {
		m[f.ID] = f
	}
	c.formats = formats
	c.index = m
	c.mu.Unlock()
}
