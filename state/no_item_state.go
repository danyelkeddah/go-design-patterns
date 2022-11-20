package state

import "fmt"

type NoItemState struct {
	VendingMachine *VendingMachine
}

func (s *NoItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (s *NoItemState) AddItem(count int) error {
	s.VendingMachine.IncrementItemCount(count)
	s.VendingMachine.SetState(s.VendingMachine.HasItem)
	return nil
}

func (s *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (s *NoItemState) DispenseItem() error {
	return fmt.Errorf("item out of stock")
}
