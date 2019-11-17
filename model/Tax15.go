package model

// Tax related to an investment fund order.
type Tax15 struct {

	// Type of tax applied.
	Type *TaxType13Code `xml:"Tp"`

	// Type of tax applied.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis2Code `xml:"Bsis,omitempty"`

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	ExtendedBasis *Extended350Code `xml:"XtndedBsis,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptReason1Code `xml:"XmptnRsn,omitempty"`

	// Reason for a tax exemption.
	ExtendedExemptionReason *Extended350Code `xml:"XtndedXmptnRsn,omitempty"`
}

func (t *Tax15) SetType(value string) {
	t.Type = (*TaxType13Code)(&value)
}

func (t *Tax15) SetExtendedType(value string) {
	t.ExtendedType = (*Extended350Code)(&value)
}

func (t *Tax15) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax15) SetBasis(value string) {
	t.Basis = (*TaxationBasis2Code)(&value)
}

func (t *Tax15) SetExtendedBasis(value string) {
	t.ExtendedBasis = (*Extended350Code)(&value)
}

func (t *Tax15) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax15) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax15) SetExemptionReason(value string) {
	t.ExemptionReason = (*TaxExemptReason1Code)(&value)
}

func (t *Tax15) SetExtendedExemptionReason(value string) {
	t.ExtendedExemptionReason = (*Extended350Code)(&value)
}
