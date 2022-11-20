package abstract_factory

import (
	"fmt"
	"github.com/danyelkeddah/go-design-patterns/abstract_factory/contracts"
)

type Adidas struct{}

func (a *Adidas) MakeShoe() contracts.Shoe {
	fmt.Println("making adidas shoe...")
	return &AdidasShoe{
		Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() contracts.Shirt {
	fmt.Println("making adidas shirt...")
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
