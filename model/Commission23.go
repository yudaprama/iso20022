package model

// Amount of money due to a party as compensation for a service.
type Commission23 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType3Choice `xml:"Tp"`

	// Basis upon which a commission is charged, for example, flat fee.
	Basis *CommissionBasis1Choice `xml:"Bsis,omitempty"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Party entitled to the amount of money resulting from a commission.
	RecipientIdentification *PartyIdentification70Choice `xml:"RcptId,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Voluntary non-enforcement of the right to part of a commission.
	WaivingDetails *CommissionWaiver4 `xml:"WvgDtls,omitempty"`
}

func (c *Commission23) AddType() *CommissionType3Choice {
	c.Type = new(CommissionType3Choice)
	return c.Type
}

func (c *Commission23) AddBasis() *CommissionBasis1Choice {
	c.Basis = new(CommissionBasis1Choice)
	return c.Basis
}

func (c *Commission23) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission23) AddRecipientIdentification() *PartyIdentification70Choice {
	c.RecipientIdentification = new(PartyIdentification70Choice)
	return c.RecipientIdentification
}

func (c *Commission23) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

func (c *Commission23) AddWaivingDetails() *CommissionWaiver4 {
	c.WaivingDetails = new(CommissionWaiver4)
	return c.WaivingDetails
}
