package ledger

import "github.com/bft/protos"

type Ledger struct {

}

func (ledger *Ledger) BeginTxBatch(id interface{}) error {
	return nil
}

func (ledger *Ledger) CommitTxBatch(id interface{}, transactions []*protos.Transaction, transactionResults []*protos.TransactionResult, metadata []byte) error {
	return nil
}
