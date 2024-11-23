package cache

// 	Implement a cache class for go with put/get methods
//and remove the oldest member when capacity is reached

type Cache[K comparable, V any] struct {
	itemsMap map[K]V
	itemsList []K
	maxCapacity int
	capacity int
}


func CreateCache[K comparable, V any](maxCapacity int) *Cache[K, V] {
	return &Cache[K, V]{
		itemsMap: make(map[K]V),
		itemsList: make([]K, 0),
		maxCapacity: maxCapacity,
		capacity: 0,
	}
}

func (c *Cache[K,V]) Put(key K, value V) {
	c.removeOldestIfNeeded()
	
	c.itemsMap[key] = value

	c.itemsList = append(c.itemsList, key)
	c.capacity++
}

func (c *Cache[K,V]) Get(key K) V {
	return c.itemsMap[key]
}

/**
	Remove the oldest item if the capacity is reached
	@return void
*/
func (c *Cache[K,V]) removeOldestIfNeeded() {
	if c.capacity == c.maxCapacity {
		oldestKey := c.itemsList[0]

		delete(c.itemsMap, oldestKey)

		c.itemsList = c.itemsList[1:]
		c.capacity--
	}
}
