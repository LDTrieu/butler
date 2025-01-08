package cache

type Cache struct {
	data map[string]any
}

func InitCache() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	c.data[key] = value
}

func (c *Cache) Get(key string) any {
	return c.data[key]
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
