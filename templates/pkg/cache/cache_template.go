package cache

import (
	"sync"
	"time"

	"github.com/spf13/viper"
)

type CKey interface{}

type CEntry struct {
	Value      interface{}
	Expiration time.Time
}

type Cache struct {
	store sync.Map
	ttl   time.Duration
}

var instance *Cache
var once sync.Once

func GetInstance() *Cache {
	once.Do(func() {
		instance = &Cache{
			ttl: viper.GetDuration("config.ttl") * time.Minute,
		}
	})
	return instance
}

func (c *Cache) Get(key CKey) (interface{}, bool) {
	if entry, ok := c.store.Load(key); ok {
		cacheEntry := entry.(*CEntry)
		if time.Now().Before(cacheEntry.Expiration) {
			return cacheEntry.Value, true
		}
		c.Clear(key)
	}
	return nil, false
}

func (c *Cache) Set(key CKey, value interface{}) {
	expiration := time.Now().Add(c.ttl)
	c.store.Store(key, &CEntry{
		Value:      value,
		Expiration: expiration,
	})
}

func (c *Cache) Clear(key CKey) {
	c.store.Delete(key)
}
