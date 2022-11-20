package facade

import "fmt"

type Ledger struct{}

func (l *Ledger) MakeEntry(accountId, txnType string, amount int) {
	fmt.Printf("Make ledger entry for account #%s with transaction type %s for amount of %d", accountId, txnType, amount)
}
