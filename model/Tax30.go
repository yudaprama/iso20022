package model

// Tax related to an investment fund order.
type Tax30 struct {

	// Type of tax.
	Type *TaxType3Choice `xml:"Tp"`

	// Tax to be applied.
	Tax *TaxAmountOrRate4Choice `xml:"Tax,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for the tax exemption.
	ExemptionReason *ExemptionReason1Choice `xml:"XmptnRsn,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification113 `xml:"RcptId,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation9 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax30) AddType() *TaxType3Choice {
	t.Type = new(TaxType3Choice)
	return t.Type
}

func (t *Tax30) AddTax() *TaxAmountOrRate4Choice {
	t.Tax = new(TaxAmountOrRate4Choice)
	return t.Tax
}

func (t *Tax30) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax30) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax30) AddExemptionReason() *ExemptionReason1Choice {
	t.ExemptionReason = new(ExemptionReason1Choice)
	return t.ExemptionReason
}

func (t *Tax30) AddRecipientIdentification() *PartyIdentification113 {
	t.RecipientIdentification = new(PartyIdentification113)
	return t.RecipientIdentification
}

func (t *Tax30) AddTaxCalculationDetails() *TaxCalculationInformation9 {
	t.TaxCalculationDetails = new(TaxCalculationInformation9)
	return t.TaxCalculationDetails
}
