package mediator

import "fmt"

type GoodsTrain struct {
	Mediator Mediator
}

func (p *GoodsTrain) RequestArrival() {
	if p.Mediator.CanLand(p) {
		fmt.Println("GoodsTrain: Landing")
	} else {
		fmt.Println("GoodsTrain: Waiting")
	}
}

func (p *GoodsTrain) Departure() {
	fmt.Println("GoodsTrain: Leaving")
	p.Mediator.NotifyFree()
}

func (p *GoodsTrain) PermitArrival() {
	fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}
