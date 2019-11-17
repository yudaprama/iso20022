package model

// Choice of format for the cancellation reason.
type CancellationReason28Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason5Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CancellationReason28Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason5Code)(&value)
}

func (c *CancellationReason28Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
