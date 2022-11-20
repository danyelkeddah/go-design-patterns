package contracts

type SportsFactory interface {
	MakeShoe() Shoe
	MakeShirt() Shirt
}
