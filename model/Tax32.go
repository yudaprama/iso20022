package model

// Tax related to an investment fund order.
type Tax32 struct {

	// Type of tax applied.
	Type *TaxType3Choice `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax. This amount is provided for information only.
	InformativeAmount *ActiveCurrencyAndAmount `xml:"InftvAmt,omitempty"`

	// Rate used to calculate the tax. This rate is provided for information only.
	InformativeRate *PercentageRate `xml:"InftvRate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for the tax exemption.
	ExemptionReason *ExemptionReason1Choice `xml:"XmptnRsn,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification113 `xml:"RcptId,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation10 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax32) AddType() *TaxType3Choice {
	t.Type = new(TaxType3Choice)
	return t.Type
}

func (t *Tax32) SetInformativeAmount(value, currency string) {
	t.InformativeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *Tax32) SetInformativeRate(value string) {
	t.InformativeRate = (*PercentageRate)(&value)
}

func (t *Tax32) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax32) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax32) AddExemptionReason() *ExemptionReason1Choice {
	t.ExemptionReason = new(ExemptionReason1Choice)
	return t.ExemptionReason
}

func (t *Tax32) AddRecipientIdentification() *PartyIdentification113 {
	t.RecipientIdentification = new(PartyIdentification113)
	return t.RecipientIdentification
}

func (t *Tax32) AddTaxCalculationDetails() *TaxCalculationInformation10 {
	t.TaxCalculationDetails = new(TaxCalculationInformation10)
	return t.TaxCalculationDetails
}
