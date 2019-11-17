package model

// Specification of the charge type.
type ChargeType1 struct {

	// Structured format.
	Structured *ChargeType6Code `xml:"Strd"`

	// Additional information about the type of charge.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ChargeType1) SetStructured(value string) {
	c.Structured = (*ChargeType6Code)(&value)
}

func (c *ChargeType1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
