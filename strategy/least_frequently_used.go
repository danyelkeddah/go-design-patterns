package strategy

import "fmt"

type LeastFrequentlyUsed struct {
}

func (l LeastFrequentlyUsed) Delete(c *Cache) {
	fmt.Println("Delete using lesat frequently used strategy")
}
