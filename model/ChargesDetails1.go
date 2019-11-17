package model

// Amount of money associated with a service.
type ChargesDetails1 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType8Code `xml:"Tp"`

	// Specifies the type of charge by means of a code.
	OtherChargesType *Max35Text `xml:"OthrChrgsTp"`

	// Amount of money asked or paid for the charge.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`
}

func (c *ChargesDetails1) SetType(value string) {
	c.Type = (*ChargeType8Code)(&value)
}

func (c *ChargesDetails1) SetOtherChargesType(value string) {
	c.OtherChargesType = (*Max35Text)(&value)
}

func (c *ChargesDetails1) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}

func (c *ChargesDetails1) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}
