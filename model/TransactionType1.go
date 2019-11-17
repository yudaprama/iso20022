package model

// Set of elements used to identify the transactions to be reported.
type TransactionType1 struct {

	// Specifies the status on the books of the account servicer of the transactions to be reported.
	Status *EntryStatus2Code `xml:"Sts"`

	// Indicates whether the reporting request refers to credit or debit entries.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies the minimum value of entries to be reported in the requested message.
	FloorLimit []*Limit2 `xml:"FlrLmt,omitempty"`
}

func (t *TransactionType1) SetStatus(value string) {
	t.Status = (*EntryStatus2Code)(&value)
}

func (t *TransactionType1) SetCreditDebitIndicator(value string) {
	t.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (t *TransactionType1) AddFloorLimit() *Limit2 {
	newValue := new(Limit2)
	t.FloorLimit = append(t.FloorLimit, newValue)
	return newValue
}
