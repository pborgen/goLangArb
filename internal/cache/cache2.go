package cache

type Cache2[K comparable, V any] struct {
	itemsMap map[K]V
	itemsList []K
	maxCapacity int
	capacity int
}

func CreateCache2[K comparable, V any](maxCapacity int) *Cache2[K, V] {
	return &Cache2[K, V]{
		itemsMap: make(map[K]V),
		itemsList: make([]K, 0),
	}
}

func (c *Cache2[K, V]) Put(key K, value V) {

	c.removeOldestIfNeeded()

	c.itemsMap[key] = value
	c.itemsList = append(c.itemsList, key)
}

func (c *Cache2[K, V]) Get(key K) V {
	return c.itemsMap[key]
}

func (c *Cache2[K, V]) removeOldestIfNeeded() {
	if c.capacity == c.maxCapacity {
		oldestKey := c.itemsList[0]

		delete(c.itemsMap, oldestKey)
		
		c.itemsList = c.itemsList[1:]
		c.capacity--
	}	
}