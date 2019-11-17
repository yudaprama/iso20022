package model

// Set of elements used to provide information on the charges related to the payment transaction.
type Charges2 struct {

	// Transaction charges to be paid by the charge bearer.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Agent that takes the transaction charges or to which the transaction charges are due.
	Agent *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

func (c *Charges2) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charges2) AddAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.Agent = new(BranchAndFinancialInstitutionIdentification5)
	return c.Agent
}
