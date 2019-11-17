package model

// Amount of money associated with a service.
type Charge4 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeTypeFormat2Choice `xml:"Tp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Method used to calculate a charge.
	ChargeBasis *TaxationBasis2Code `xml:"ChrgBsis,omitempty"`

	// Specifies the party that will bear the charges associated with a transfer.
	ChargeBearer *ChargeBearer1Code `xml:"ChrgBr,omitempty"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification1Choice `xml:"RcptId,omitempty"`
}

func (c *Charge4) AddType() *ChargeTypeFormat2Choice {
	c.Type = new(ChargeTypeFormat2Choice)
	return c.Type
}

func (c *Charge4) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charge4) SetChargeBasis(value string) {
	c.ChargeBasis = (*TaxationBasis2Code)(&value)
}

func (c *Charge4) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearer1Code)(&value)
}

func (c *Charge4) AddRecipientIdentification() *PartyIdentification1Choice {
	c.RecipientIdentification = new(PartyIdentification1Choice)
	return c.RecipientIdentification
}
