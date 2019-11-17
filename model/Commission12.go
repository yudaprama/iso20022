package model

// Amount of money due to a party as compensation for a service.
type Commission12 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType7Code `xml:"Tp"`

	// Service for which the commission is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Basis upon which a commission is charged, eg, flat fee.
	Basis *TaxationBasis4Code `xml:"Bsis,omitempty"`

	// Basis upon which a commission is charged, eg, flat fee.
	ExtendedBasis *Extended350Code `xml:"XtndedBsis,omitempty"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Party entitled to the amount of money resulting from a commission.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`
}

func (c *Commission12) SetType(value string) {
	c.Type = (*CommissionType7Code)(&value)
}

func (c *Commission12) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Commission12) SetBasis(value string) {
	c.Basis = (*TaxationBasis4Code)(&value)
}

func (c *Commission12) SetExtendedBasis(value string) {
	c.ExtendedBasis = (*Extended350Code)(&value)
}

func (c *Commission12) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission12) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

func (c *Commission12) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}
