package mediator

import "fmt"

type PassengerTrain struct {
	Mediator Mediator
}

func (p *PassengerTrain) RequestArrival() {
	if p.Mediator.CanLand(p) {
		fmt.Println("PassengerTrain: Landing")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (p *PassengerTrain) Departure() {
	fmt.Println("PassengerTrain: Leaving")
	p.Mediator.NotifyFree()
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}
