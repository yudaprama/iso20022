package model

// Choice of format for the cancellation reason.
type CancellationReason20Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason12Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CancellationReason20Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason12Code)(&value)
}

func (c *CancellationReason20Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
