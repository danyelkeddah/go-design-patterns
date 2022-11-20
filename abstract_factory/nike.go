package abstract_factory

import (
	"fmt"
	"github.com/danyelkeddah/go-design-patterns/abstract_factory/contracts"
)

type Nike struct{}

func (n *Nike) MakeShoe() contracts.Shoe {
	fmt.Println("making nike shoe...")
	return &NikeShoe{
		Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (n *Nike) MakeShirt() contracts.Shirt {
	fmt.Println("making nike shirt...")
	return &NikeShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
