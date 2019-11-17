package model

// Provides further individual record details on the charges related to the payment transaction.
type ChargesRecord1 struct {

	// Transaction charges to be paid by the charge bearer.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the charges amount is a credit or a debit amount.
	// Usage: A zero amount is considered to be a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Specifies the type of charge.
	Type *ChargeType3Choice `xml:"Tp,omitempty"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	Bearer *ChargeBearerType1Code `xml:"Br,omitempty"`

	// Agent that takes the transaction charges or to which the transaction charges are due.
	Agent *BranchAndFinancialInstitutionIdentification5 `xml:"Agt,omitempty"`

	// Provides details on the tax applied to charges.
	Tax *TaxCharges2 `xml:"Tax,omitempty"`
}

func (c *ChargesRecord1) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *ChargesRecord1) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *ChargesRecord1) AddType() *ChargeType3Choice {
	c.Type = new(ChargeType3Choice)
	return c.Type
}

func (c *ChargesRecord1) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *ChargesRecord1) SetBearer(value string) {
	c.Bearer = (*ChargeBearerType1Code)(&value)
}

func (c *ChargesRecord1) AddAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.Agent = new(BranchAndFinancialInstitutionIdentification5)
	return c.Agent
}

func (c *ChargesRecord1) AddTax() *TaxCharges2 {
	c.Tax = new(TaxCharges2)
	return c.Tax
}
