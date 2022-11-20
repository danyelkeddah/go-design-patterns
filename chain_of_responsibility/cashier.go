package chain_of_responsibility

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment already done")
		return
	}
	fmt.Println("Cashier getting money from the patient")
}
func (c *Cashier) SetNext(next Department) {
	c.next = next
}
