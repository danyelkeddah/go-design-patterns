package decorator

type PeppyPaneer struct{}

func (p *PeppyPaneer) GetPrice() int {
	return 20
}
