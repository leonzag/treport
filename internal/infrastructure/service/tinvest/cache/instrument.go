package cache

import (
	"sync"

	"github.com/leonzag/treport/internal/domain/entity"
)

type uidInstr map[string]*entity.Instrument

type InstrumentCache struct {
	cache uidInstr
	mu    sync.RWMutex
}

func NewInstrumentCache() *InstrumentCache {
	return &InstrumentCache{
		cache: make(uidInstr),
	}
}

func (c *InstrumentCache) Add(uid string, i *entity.Instrument) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[uid] = i
}

func (c *InstrumentCache) Get(uid string) (*entity.Instrument, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	i, ok := c.cache[uid]
	return i, ok
}

func (c *InstrumentCache) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(uidInstr)
}
