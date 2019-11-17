package model

// Tax related to an investment fund order.
type Tax7 struct {

	// Type of tax applied.
	Type *TaxType2 `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptionReason1 `xml:"XmptnRsn,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation3 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax7) AddType() *TaxType2 {
	t.Type = new(TaxType2)
	return t.Type
}

func (t *Tax7) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax7) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *Tax7) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax7) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax7) AddExemptionReason() *TaxExemptionReason1 {
	t.ExemptionReason = new(TaxExemptionReason1)
	return t.ExemptionReason
}

func (t *Tax7) AddTaxCalculationDetails() *TaxCalculationInformation3 {
	t.TaxCalculationDetails = new(TaxCalculationInformation3)
	return t.TaxCalculationDetails
}
