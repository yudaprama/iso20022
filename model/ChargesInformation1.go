package model

// Information on the charges related to the payment transaction.
type ChargesInformation1 struct {

	// Transaction charges to be paid by the charge bearer.
	ChargesAmount *CurrencyAndAmount `xml:"ChrgsAmt"`

	// Party that takes the transaction charges or to which the transaction charges are due.
	ChargesParty *BranchAndFinancialInstitutionIdentification3 `xml:"ChrgsPty"`
}

func (c *ChargesInformation1) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation1) AddChargesParty() *BranchAndFinancialInstitutionIdentification3 {
	c.ChargesParty = new(BranchAndFinancialInstitutionIdentification3)
	return c.ChargesParty
}
