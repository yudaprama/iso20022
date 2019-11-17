package model

// Set of elements used to provide information on the charges related to the payment transaction.
type ChargesInformation5 struct {

	// Transaction charges to be paid by the charge bearer.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Party that takes the transaction charges or to which the transaction charges are due.
	Party *BranchAndFinancialInstitutionIdentification4 `xml:"Pty"`
}

func (c *ChargesInformation5) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation5) AddParty() *BranchAndFinancialInstitutionIdentification4 {
	c.Party = new(BranchAndFinancialInstitutionIdentification4)
	return c.Party
}
