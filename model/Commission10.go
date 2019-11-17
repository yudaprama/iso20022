package model

// Amount of money due to a party as compensation for a service.
type Commission10 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType6Code `xml:"Tp,omitempty"`

	// Service for which the commission is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`

	// Basis upon which a commission is charged, eg, flat fee.
	Basis *TaxationBasis4Code `xml:"Bsis,omitempty"`

	// Basis upon which a commission is charged, eg, flat fee.
	ExtendedBasis *Extended350Code `xml:"XtndedBsis,omitempty"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Party entitled to the amount of money resulting from a commission.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Voluntary non-enforcement of the right to all or part of a commission.
	WaivingDetails *CommissionWaiver3 `xml:"WvgDtls,omitempty"`
}

func (c *Commission10) SetType(value string) {
	c.Type = (*CommissionType6Code)(&value)
}

func (c *Commission10) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Commission10) SetBasis(value string) {
	c.Basis = (*TaxationBasis4Code)(&value)
}

func (c *Commission10) SetExtendedBasis(value string) {
	c.ExtendedBasis = (*Extended350Code)(&value)
}

func (c *Commission10) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission10) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Commission10) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

func (c *Commission10) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

func (c *Commission10) AddWaivingDetails() *CommissionWaiver3 {
	c.WaivingDetails = new(CommissionWaiver3)
	return c.WaivingDetails
}
