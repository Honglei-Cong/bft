package ledger

import "sync"

type LedgerManager struct {

}

var once sync.Once

func GetLedger() (*Ledger, error) {
	once.Do(func() {

	})
	return nil, nil
}