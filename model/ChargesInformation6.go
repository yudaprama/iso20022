package model

// Set of elements used to provide information on the charges related to the payment transaction.
type ChargesInformation6 struct {

	// Total of all charges and taxes applied to the entry.
	TotalChargesAndTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty"`

	// Transaction charges to be paid by the charge bearer.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the charges amount is a credit or a debit amount.
	// Usage: A zero amount is considered to be a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Specifies the type of charge.
	Type *ChargeType2Choice `xml:"Tp,omitempty"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	Bearer *ChargeBearerType1Code `xml:"Br,omitempty"`

	// Party that takes the transaction charges or to which the transaction charges are due.
	Party *BranchAndFinancialInstitutionIdentification4 `xml:"Pty,omitempty"`

	// Set of elements used to provide details on the tax applied to charges.
	Tax *TaxCharges2 `xml:"Tax,omitempty"`
}

func (c *ChargesInformation6) SetTotalChargesAndTaxAmount(value, currency string) {
	c.TotalChargesAndTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation6) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation6) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *ChargesInformation6) AddType() *ChargeType2Choice {
	c.Type = new(ChargeType2Choice)
	return c.Type
}

func (c *ChargesInformation6) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *ChargesInformation6) SetBearer(value string) {
	c.Bearer = (*ChargeBearerType1Code)(&value)
}

func (c *ChargesInformation6) AddParty() *BranchAndFinancialInstitutionIdentification4 {
	c.Party = new(BranchAndFinancialInstitutionIdentification4)
	return c.Party
}

func (c *ChargesInformation6) AddTax() *TaxCharges2 {
	c.Tax = new(TaxCharges2)
	return c.Tax
}
