package model

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type Tax3 struct {

	// Type of tax applied.
	Type *TaxTypeFormat2Choice `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis2Code `xml:"Bsis,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification1Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptionReasonFormatChoice `xml:"XmptnRsn,omitempty"`
}

func (t *Tax3) AddType() *TaxTypeFormat2Choice {
	t.Type = new(TaxTypeFormat2Choice)
	return t.Type
}

func (t *Tax3) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax3) SetBasis(value string) {
	t.Basis = (*TaxationBasis2Code)(&value)
}

func (t *Tax3) AddRecipientIdentification() *PartyIdentification1Choice {
	t.RecipientIdentification = new(PartyIdentification1Choice)
	return t.RecipientIdentification
}

func (t *Tax3) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax3) AddExemptionReason() *TaxExemptionReasonFormatChoice {
	t.ExemptionReason = new(TaxExemptionReasonFormatChoice)
	return t.ExemptionReason
}
