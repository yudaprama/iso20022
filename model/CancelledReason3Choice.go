package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request is cancelled.
type CancelledReason3Choice struct {

	// Standard code to specify the reason why the instruction is cancelled.
	Code *CancelledStatusReason11Code `xml:"Cd"`

	// Proprietary identification specifying the reason why the instruction is cancelled.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CancelledReason3Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason11Code)(&value)
}

func (c *CancelledReason3Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
