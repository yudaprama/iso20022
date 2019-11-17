package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request is cancelled.
type CancelledReason11Choice struct {

	// Standard code to specify the reason why the instruction is cancelled.
	Code *CancelledStatusReason6Code `xml:"Cd"`

	// Proprietary identification specifying the reason why the instruction is cancelled.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CancelledReason11Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason6Code)(&value)
}

func (c *CancelledReason11Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
