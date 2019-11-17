package model

// Identifies the underlying transaction(s) and/or batched entries.
type EntryDetails3 struct {

	// Provides details on batched transactions.
	Batch *BatchInformation2 `xml:"Btch,omitempty"`

	// Provides information on the underlying transaction(s).
	TransactionDetails []*EntryTransaction4 `xml:"TxDtls,omitempty"`
}

func (e *EntryDetails3) AddBatch() *BatchInformation2 {
	e.Batch = new(BatchInformation2)
	return e.Batch
}

func (e *EntryDetails3) AddTransactionDetails() *EntryTransaction4 {
	newValue := new(EntryTransaction4)
	e.TransactionDetails = append(e.TransactionDetails, newValue)
	return newValue
}
