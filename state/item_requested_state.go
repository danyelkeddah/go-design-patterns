package state

import "fmt"

type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (s *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requeted")
}
func (s *ItemRequestedState) AddItem(money int) error {
	return fmt.Errorf("item dispense in progress")
}

func (s *ItemRequestedState) InsertMoney(money int) error {
	if money < s.VendingMachine.ItemPrice {
		return fmt.Errorf("inserted money is less. please insert %d", s.VendingMachine.ItemPrice)
	}
	fmt.Println("money entered is ok")
	s.VendingMachine.SetState(s.VendingMachine.HasMoney)
	return nil
}

func (s *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}
