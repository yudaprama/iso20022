package model

// Choice of format for the cancellation reason.
type CancellationReason19Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason13Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CancellationReason19Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason13Code)(&value)
}

func (c *CancellationReason19Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
