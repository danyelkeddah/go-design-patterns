package strategy

type DeletionAlgo interface {
	Delete(c *Cache)
}
