package model

// Tax related to an investment fund order.
type Tax14 struct {

	// Type of tax applied.
	Type *TaxType11Code `xml:"Tp"`

	// Type of tax applied.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptReason1Code `xml:"XmptnRsn,omitempty"`

	// Reason for a tax exemption.
	ExtendedExemptionReason *Extended350Code `xml:"XtndedXmptnRsn,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation6 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax14) SetType(value string) {
	t.Type = (*TaxType11Code)(&value)
}

func (t *Tax14) SetExtendedType(value string) {
	t.ExtendedType = (*Extended350Code)(&value)
}

func (t *Tax14) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax14) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *Tax14) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax14) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax14) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax14) SetExemptionReason(value string) {
	t.ExemptionReason = (*TaxExemptReason1Code)(&value)
}

func (t *Tax14) SetExtendedExemptionReason(value string) {
	t.ExtendedExemptionReason = (*Extended350Code)(&value)
}

func (t *Tax14) AddTaxCalculationDetails() *TaxCalculationInformation6 {
	t.TaxCalculationDetails = new(TaxCalculationInformation6)
	return t.TaxCalculationDetails
}
