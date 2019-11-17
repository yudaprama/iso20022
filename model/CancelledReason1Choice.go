package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request is cancelled.
type CancelledReason1Choice struct {

	// Standard code to specify the reason why the instruction is cancelled.
	Code *CancelledStatusReason6Code `xml:"Cd"`

	// Proprietary identification specifying the reason why the instruction is cancelled.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CancelledReason1Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason6Code)(&value)
}

func (c *CancelledReason1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
