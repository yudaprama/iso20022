package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request is cancelled.
type CancelledReason9Choice struct {

	// Standard code to specify the reason why the instruction is cancelled.
	Code *CancelledStatusReason5Code `xml:"Cd"`

	// Proprietary identification specifying the reason why the instruction is cancelled.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CancelledReason9Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason5Code)(&value)
}

func (c *CancelledReason9Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
