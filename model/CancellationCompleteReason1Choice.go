package model

// Choice of formats for the cancelled complete reason.
type CancellationCompleteReason1Choice struct {

	// Cancelled complete reason expressed as a code.
	Code *CancelledStatusReason1Code `xml:"Cd"`

	// Cancelled complete reason expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (c *CancellationCompleteReason1Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason1Code)(&value)
}

func (c *CancellationCompleteReason1Choice) AddProprietary() *GenericIdentification36 {
	c.Proprietary = new(GenericIdentification36)
	return c.Proprietary
}
