package strategy

import "fmt"

type LeastRecentlyUsed struct{}

func (l LeastRecentlyUsed) Delete(c *Cache) {
	fmt.Println("Delete using least recently used strategy")
}
