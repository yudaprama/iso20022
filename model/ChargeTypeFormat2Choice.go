package model

// Type of charge expressed either in free format or in structured or free form.
type ChargeTypeFormat2Choice struct {

	// Specifies the type of charge in free text form.
	Unstructured *Max35Text `xml:"Ustrd"`

	// Specifies the type of charge in a structured form.
	Structured *ChargeType4Code `xml:"Strd"`
}

func (c *ChargeTypeFormat2Choice) SetUnstructured(value string) {
	c.Unstructured = (*Max35Text)(&value)
}

func (c *ChargeTypeFormat2Choice) SetStructured(value string) {
	c.Structured = (*ChargeType4Code)(&value)
}
