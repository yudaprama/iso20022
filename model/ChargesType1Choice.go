package model

// Specifies the type of charges as a code or free text.
type ChargesType1Choice struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType8Code `xml:"Tp"`

	// Specifies the type of charge by means of a free text
	OtherChargesType *Max35Text `xml:"OthrChrgsTp"`
}

func (c *ChargesType1Choice) SetType(value string) {
	c.Type = (*ChargeType8Code)(&value)
}

func (c *ChargesType1Choice) SetOtherChargesType(value string) {
	c.OtherChargesType = (*Max35Text)(&value)
}
