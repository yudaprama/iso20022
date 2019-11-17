package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request is cancelled.
type CancelledReason10Choice struct {

	// Standard code to specify the reason why the instruction is cancelled.
	Code *CancelledStatusReason5Code `xml:"Cd"`

	// Proprietary identification specifying the reason why the instruction is cancelled.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CancelledReason10Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason5Code)(&value)
}

func (c *CancelledReason10Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
