package state

import "fmt"

type HasItemState struct {
	VendingMachine *VendingMachine
}

func (s *HasItemState) RequestItem() error {
	if s.VendingMachine.ItemCount == 0 {
		s.VendingMachine.SetState(s.VendingMachine.NoItem)
		return fmt.Errorf("no item present")
	}
	fmt.Println("item requested")
	s.VendingMachine.SetState(s.VendingMachine.ItemRequested)
	return nil
}

func (s *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	s.VendingMachine.IncrementItemCount(count)
	return nil
}

func (s *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (s *HasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}
