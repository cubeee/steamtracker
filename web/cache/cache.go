package cache

var (
	GlobalCache *Cache
)

type Cache struct {
	indexCache IndexCache
}

func (c *Cache) SetIndexCache(cache IndexCache) {
	c.indexCache = cache
}

func (c *Cache) GetIndexCache() IndexCache {
	return c.indexCache
}
