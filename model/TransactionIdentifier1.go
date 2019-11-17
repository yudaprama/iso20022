package model

// Identification of the transaction in an unambiguous way.
type TransactionIdentifier1 struct {

	// Local date and time of the transaction assigned by the POI (Point Of Interaction).
	TransactionDateTime *ISODateTime `xml:"TxDtTm"`

	// Identification of the transaction that has to be unique for a time period.
	TransactionReference *Max35Text `xml:"TxRef"`
}

func (t *TransactionIdentifier1) SetTransactionDateTime(value string) {
	t.TransactionDateTime = (*ISODateTime)(&value)
}

func (t *TransactionIdentifier1) SetTransactionReference(value string) {
	t.TransactionReference = (*Max35Text)(&value)
}
