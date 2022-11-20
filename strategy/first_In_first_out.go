package strategy

import "fmt"

type FirstInFirstOut struct {
}

func (s FirstInFirstOut) Delete(c *Cache) {
	fmt.Println("Delete using first in first out strategy")
}
