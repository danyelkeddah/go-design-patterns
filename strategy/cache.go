package strategy

type Cache struct {
	Storage      map[string]string
	DeletionAlgo DeletionAlgo
	Capacity     int
	MaxCapacity  int
}

func NewCache(deletionAlgo DeletionAlgo) *Cache {
	return &Cache{
		Storage:      make(map[string]string),
		DeletionAlgo: deletionAlgo,
		Capacity:     0,
		MaxCapacity:  2,
	}
}

func (c *Cache) SetDeletionAlgorithm(da DeletionAlgo) {
	c.DeletionAlgo = da
}

func (c *Cache) Add(key, value string) {
	if c.Capacity == c.MaxCapacity {
		c.Clean()
	}
	c.Capacity++
	c.Storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.Storage, key)
}

func (c *Cache) Clean() {
	c.DeletionAlgo.Delete(c)
	c.Capacity--
}
