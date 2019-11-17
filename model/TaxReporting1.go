package model

// Information for tax reporting.
type TaxReporting1 struct {

	// Country of taxation of the organisation or individual person.
	TaxationCountry *CountryCode `xml:"TaxtnCtry"`

	// Tax rate to be applied.
	TaxRate *PercentageRate `xml:"TaxRate,omitempty"`

	// Party that pays the tax.
	TaxPayer *PartyIdentification70Choice `xml:"TaxPyer,omitempty"`

	// Party that receives the tax.
	TaxRecipient *PartyIdentification70Choice `xml:"TaxRcpt,omitempty"`

	// Cash account information for the payment of tax.
	CashAccountDetails *CashAccount32 `xml:"CshAcctDtls,omitempty"`

	// Additional information for the reporting of tax.
	Description *Max350Text `xml:"Desc,omitempty"`
}

func (t *TaxReporting1) SetTaxationCountry(value string) {
	t.TaxationCountry = (*CountryCode)(&value)
}

func (t *TaxReporting1) SetTaxRate(value string) {
	t.TaxRate = (*PercentageRate)(&value)
}

func (t *TaxReporting1) AddTaxPayer() *PartyIdentification70Choice {
	t.TaxPayer = new(PartyIdentification70Choice)
	return t.TaxPayer
}

func (t *TaxReporting1) AddTaxRecipient() *PartyIdentification70Choice {
	t.TaxRecipient = new(PartyIdentification70Choice)
	return t.TaxRecipient
}

func (t *TaxReporting1) AddCashAccountDetails() *CashAccount32 {
	t.CashAccountDetails = new(CashAccount32)
	return t.CashAccountDetails
}

func (t *TaxReporting1) SetDescription(value string) {
	t.Description = (*Max350Text)(&value)
}
