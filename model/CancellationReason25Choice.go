package model

// Choice of format for the cancellation reason.
type CancellationReason25Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason9Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CancellationReason25Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason9Code)(&value)
}

func (c *CancellationReason25Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
