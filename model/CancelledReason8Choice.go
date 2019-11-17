package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request is cancelled.
type CancelledReason8Choice struct {

	// Standard code to specify the reason why the instruction is cancelled.
	Code *CancelledStatusReason6Code `xml:"Cd"`

	// Proprietary identification specifying the reason why the instruction is cancelled.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CancelledReason8Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason6Code)(&value)
}

func (c *CancelledReason8Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
