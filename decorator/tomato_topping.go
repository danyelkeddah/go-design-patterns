package decorator

type TomatoTopping struct {
	Pizza Pizza
}

func (t *TomatoTopping) GetPrice() int {
	pizzaPrize := t.Pizza.GetPrice()
	return pizzaPrize + 7
}
