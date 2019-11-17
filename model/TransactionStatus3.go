package model

// Identifies the status of the transaction by means of a code.
type TransactionStatus3 struct {

	// Identifies the status of the transaction by means of a code.
	Status *BaselineStatus2Code `xml:"Sts"`
}

func (t *TransactionStatus3) SetStatus(value string) {
	t.Status = (*BaselineStatus2Code)(&value)
}
