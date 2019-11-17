package model

// Amount of money associated with a service.
type Charge8 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType1 `xml:"Tp"`

	// Method used to calculate a charge.
	ChargeBasis *TaxationBasis1 `xml:"ChrgBsis,omitempty"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`
}

func (c *Charge8) AddType() *ChargeType1 {
	c.Type = new(ChargeType1)
	return c.Type
}

func (c *Charge8) AddChargeBasis() *TaxationBasis1 {
	c.ChargeBasis = new(TaxationBasis1)
	return c.ChargeBasis
}

func (c *Charge8) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge8) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Charge8) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}
