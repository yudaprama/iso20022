package model

// Identifies the underlying transaction(s) and/or batched entries.
type EntryDetails2 struct {

	// Provides details on batched transactions.
	Batch *BatchInformation2 `xml:"Btch,omitempty"`

	// Provides information on the underlying transaction(s).
	TransactionDetails []*EntryTransaction3 `xml:"TxDtls,omitempty"`
}

func (e *EntryDetails2) AddBatch() *BatchInformation2 {
	e.Batch = new(BatchInformation2)
	return e.Batch
}

func (e *EntryDetails2) AddTransactionDetails() *EntryTransaction3 {
	newValue := new(EntryTransaction3)
	e.TransactionDetails = append(e.TransactionDetails, newValue)
	return newValue
}
