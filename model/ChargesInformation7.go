package model

// Set of elements used to provide information on the charges related to the payment transaction.
type ChargesInformation7 struct {

	// Transaction charges to be paid by the charge bearer.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Party that takes the transaction charges or to which the transaction charges are due.
	Party *BranchAndFinancialInstitutionIdentification5 `xml:"Pty"`
}

func (c *ChargesInformation7) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation7) AddParty() *BranchAndFinancialInstitutionIdentification5 {
	c.Party = new(BranchAndFinancialInstitutionIdentification5)
	return c.Party
}
