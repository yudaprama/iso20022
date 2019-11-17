package model

// Choice of format for the cancellation reason.
type CancellationReason23Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason9Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CancellationReason23Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason9Code)(&value)
}

func (c *CancellationReason23Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
