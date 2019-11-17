package model

// Choice of format for the cancellation reason.
type CancellationReason32Choice struct {

	// Cancellation reason expressed as a code.
	Code *Max35Text `xml:"Cd"`

	// Cancellation reason expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CancellationReason32Choice) SetCode(value string) {
	c.Code = (*Max35Text)(&value)
}

func (c *CancellationReason32Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
