package abstract_factory

import (
	"fmt"
	"github.com/danyelkeddah/go-design-patterns/abstract_factory/contracts"
)

func MakeSportsFactory(b string) (contracts.SportsFactory, error) {
	if b == "adidas" {
		return &Adidas{}, nil
	}

	if b == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed (%s)", b)
}
