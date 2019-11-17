package model

// Identifies the status of the transaction by means of a code.
type TransactionStatus4 struct {

	// Identifies the status of the transaction by means of a code.
	Status *BaselineStatus3Code `xml:"Sts"`
}

func (t *TransactionStatus4) SetStatus(value string) {
	t.Status = (*BaselineStatus3Code)(&value)
}
