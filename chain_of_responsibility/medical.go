package chain_of_responsibility

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.MedicineDone = true
	m.next.Execute(p)
}
func (m *Medical) SetNext(next Department) {
	m.next = next
}
