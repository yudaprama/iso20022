package model

// Tax related to an investment fund order.
type Tax31 struct {

	// Type of tax.
	Type *TaxType3Choice `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	AppliedAmount *ActiveCurrencyAndAmount `xml:"ApldAmt"`

	// Rate used to calculate the tax.
	AppliedRate *PercentageRate `xml:"ApldRate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification113 `xml:"RcptId,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation10 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax31) AddType() *TaxType3Choice {
	t.Type = new(TaxType3Choice)
	return t.Type
}

func (t *Tax31) SetAppliedAmount(value, currency string) {
	t.AppliedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *Tax31) SetAppliedRate(value string) {
	t.AppliedRate = (*PercentageRate)(&value)
}

func (t *Tax31) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax31) AddRecipientIdentification() *PartyIdentification113 {
	t.RecipientIdentification = new(PartyIdentification113)
	return t.RecipientIdentification
}

func (t *Tax31) AddTaxCalculationDetails() *TaxCalculationInformation10 {
	t.TaxCalculationDetails = new(TaxCalculationInformation10)
	return t.TaxCalculationDetails
}
