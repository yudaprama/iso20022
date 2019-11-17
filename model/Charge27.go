package model

// Amount of money associated with a service.
type Charge27 struct {

	// Type of charge.
	Type *ChargeType4Choice `xml:"Tp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Method used to calculate the charge.
	ChargeBasis *ChargeBasisType1Choice `xml:"ChrgBsis,omitempty"`

	// Specifies the party that will bear the charges associated with a transfer.
	ChargeBearer *ChargeBearer1Code `xml:"ChrgBr,omitempty"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`
}

func (c *Charge27) AddType() *ChargeType4Choice {
	c.Type = new(ChargeType4Choice)
	return c.Type
}

func (c *Charge27) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charge27) AddChargeBasis() *ChargeBasisType1Choice {
	c.ChargeBasis = new(ChargeBasisType1Choice)
	return c.ChargeBasis
}

func (c *Charge27) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearer1Code)(&value)
}

func (c *Charge27) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}
