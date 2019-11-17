package model

// Amount of money due to a party as compensation for a service.
type Commission11 struct {

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Service for which the commission is asked or paid.
	Type *CommissionType6Code `xml:"Tp,omitempty"`

	// Service for which the commission is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Indicates if the CommercialAgreementReference is a new reference or not.
	NewCommercialAgreementReferenceIndicator *YesNoIndicator `xml:"NewComrclAgrmtRefInd,omitempty"`
}

func (c *Commission11) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission11) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Commission11) SetType(value string) {
	c.Type = (*CommissionType6Code)(&value)
}

func (c *Commission11) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Commission11) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

func (c *Commission11) SetNewCommercialAgreementReferenceIndicator(value string) {
	c.NewCommercialAgreementReferenceIndicator = (*YesNoIndicator)(&value)
}
