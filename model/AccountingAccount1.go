package model

// Specific trade account for recording debits and credits to general accounting, cost accounting or budget accounting.
type AccountingAccount1 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Template describing the mask of the structure for the format of the accounting account identifier for example AABBBBCC where AA represents the country, BBBB the service classification, CC the sales area.
	CostReferencePattern *Max35Text `xml:"CostRefPttrn,omitempty"`
}

func (a *AccountingAccount1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountingAccount1) SetCostReferencePattern(value string) {
	a.CostReferencePattern = (*Max35Text)(&value)
}
