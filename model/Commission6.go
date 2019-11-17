package model

// Amount of money due to a party as compensation for a service.
type Commission6 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType1 `xml:"Tp,omitempty"`

	// Basis upon which a commission is charged, eg, flat fee.
	Basis *TaxationBasis1 `xml:"Bsis,omitempty"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Party entitled to the amount of money resulting from a commission.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Voluntary non-enforcement of the right to all or part of a commission.
	WaivingDetails *CommissionWaiver2 `xml:"WvgDtls,omitempty"`
}

func (c *Commission6) AddType() *CommissionType1 {
	c.Type = new(CommissionType1)
	return c.Type
}

func (c *Commission6) AddBasis() *TaxationBasis1 {
	c.Basis = new(TaxationBasis1)
	return c.Basis
}

func (c *Commission6) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission6) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Commission6) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

func (c *Commission6) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

func (c *Commission6) AddWaivingDetails() *CommissionWaiver2 {
	c.WaivingDetails = new(CommissionWaiver2)
	return c.WaivingDetails
}
