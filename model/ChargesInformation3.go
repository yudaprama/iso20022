package model

// Information on the charges related to the payment transaction.
type ChargesInformation3 struct {

	// Total of all charges and taxes applied to the entry.
	TotalChargesAndTaxAmount *CurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty"`

	// Transaction charges to be paid by the charge bearer.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Identifies the type of charge.
	Type *ChargeTypeChoice `xml:"Tp,omitempty"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	Bearer *ChargeBearerType1Code `xml:"Br,omitempty"`

	// Party that takes the transaction charges or to which the transaction charges are due.
	Party *BranchAndFinancialInstitutionIdentification3 `xml:"Pty,omitempty"`

	// Specifies tax details applied to charges.
	Tax *TaxCharges1 `xml:"Tax,omitempty"`
}

func (c *ChargesInformation3) SetTotalChargesAndTaxAmount(value, currency string) {
	c.TotalChargesAndTaxAmount = NewCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation3) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation3) AddType() *ChargeTypeChoice {
	c.Type = new(ChargeTypeChoice)
	return c.Type
}

func (c *ChargesInformation3) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *ChargesInformation3) SetBearer(value string) {
	c.Bearer = (*ChargeBearerType1Code)(&value)
}

func (c *ChargesInformation3) AddParty() *BranchAndFinancialInstitutionIdentification3 {
	c.Party = new(BranchAndFinancialInstitutionIdentification3)
	return c.Party
}

func (c *ChargesInformation3) AddTax() *TaxCharges1 {
	c.Tax = new(TaxCharges1)
	return c.Tax
}
