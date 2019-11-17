package model

// Transaction identification.
type TransactiontIdentification4 struct {

	// Unambiguous identification of the transaction as known by the instructing party.
	TransactionIdentification *Max35Text `xml:"TxId"`
}

func (t *TransactiontIdentification4) SetTransactionIdentification(value string) {
	t.TransactionIdentification = (*Max35Text)(&value)
}
