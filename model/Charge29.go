package model

// Amount of money associated with a service.
type Charge29 struct {

	// Type of charge.
	Type *ChargeType4Choice `xml:"Tp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Method used to calculate the charge.
	ChargeBasis *ChargeBasisType1Choice `xml:"ChrgBsis,omitempty"`

	// Specifies the party that will bear the charges associated with a transfer.
	ChargeBearer *ChargeBearer1Code `xml:"ChrgBr,omitempty"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification70Choice `xml:"RcptId,omitempty"`
}

func (c *Charge29) AddType() *ChargeType4Choice {
	c.Type = new(ChargeType4Choice)
	return c.Type
}

func (c *Charge29) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charge29) AddChargeBasis() *ChargeBasisType1Choice {
	c.ChargeBasis = new(ChargeBasisType1Choice)
	return c.ChargeBasis
}

func (c *Charge29) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearer1Code)(&value)
}

func (c *Charge29) AddRecipientIdentification() *PartyIdentification70Choice {
	c.RecipientIdentification = new(PartyIdentification70Choice)
	return c.RecipientIdentification
}
