package model

// Set of elements used to identify the underlying transaction(s) and/or batched entries.
type EntryDetails1 struct {

	// Set of elements used to provide details on batched transactions.
	Batch *BatchInformation2 `xml:"Btch,omitempty"`

	// Set of elements used to provide information on the underlying transaction(s).
	TransactionDetails []*EntryTransaction2 `xml:"TxDtls,omitempty"`
}

func (e *EntryDetails1) AddBatch() *BatchInformation2 {
	e.Batch = new(BatchInformation2)
	return e.Batch
}

func (e *EntryDetails1) AddTransactionDetails() *EntryTransaction2 {
	newValue := new(EntryTransaction2)
	e.TransactionDetails = append(e.TransactionDetails, newValue)
	return newValue
}
