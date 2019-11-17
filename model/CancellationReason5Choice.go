package model

// Choice of format for the cancellation reason.
type CancellationReason5Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason9Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CancellationReason5Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason9Code)(&value)
}

func (c *CancellationReason5Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
