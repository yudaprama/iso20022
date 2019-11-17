package model

// Amount of money associated with a service.
type Charge20 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType12Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Method used to calculate a charge.
	ChargeBasis *TaxationBasis2Code `xml:"ChrgBsis,omitempty"`

	// Method used to calculate a charge.
	ExtendedChargeBasis *Extended350Code `xml:"XtndedChrgBsis,omitempty"`

	// Specifies the party that will bear the charges associated with a transfer.
	ChargeBearer *ChargeBearer1Code `xml:"ChrgBr,omitempty"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`
}

func (c *Charge20) SetType(value string) {
	c.Type = (*ChargeType12Code)(&value)
}

func (c *Charge20) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge20) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charge20) SetChargeBasis(value string) {
	c.ChargeBasis = (*TaxationBasis2Code)(&value)
}

func (c *Charge20) SetExtendedChargeBasis(value string) {
	c.ExtendedChargeBasis = (*Extended350Code)(&value)
}

func (c *Charge20) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearer1Code)(&value)
}

func (c *Charge20) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}
