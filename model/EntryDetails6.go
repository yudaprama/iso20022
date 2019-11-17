package model

// Identifies the underlying transaction(s) and/or batched entries.
type EntryDetails6 struct {

	// Provides details on batched transactions.
	Batch *BatchInformation2 `xml:"Btch,omitempty"`

	// Provides information on the underlying transaction(s).
	TransactionDetails []*EntryTransaction7 `xml:"TxDtls,omitempty"`
}

func (e *EntryDetails6) AddBatch() *BatchInformation2 {
	e.Batch = new(BatchInformation2)
	return e.Batch
}

func (e *EntryDetails6) AddTransactionDetails() *EntryTransaction7 {
	newValue := new(EntryTransaction7)
	e.TransactionDetails = append(e.TransactionDetails, newValue)
	return newValue
}
