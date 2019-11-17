package model

// Choice of format for the cancellation reason.
type CancellationReason3Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason5Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CancellationReason3Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason5Code)(&value)
}

func (c *CancellationReason3Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
