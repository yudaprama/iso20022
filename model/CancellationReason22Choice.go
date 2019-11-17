package model

// Choice of format for the cancellation reason.
type CancellationReason22Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason14Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CancellationReason22Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason14Code)(&value)
}

func (c *CancellationReason22Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
