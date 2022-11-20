package facade

import "fmt"

type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{name: accountName}
}

func (a *Account) CheckAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account name is incorrect")
	}
	fmt.Println("Account verified")
	return nil
}
