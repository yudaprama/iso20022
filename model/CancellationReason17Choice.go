package model

// Choice of format for the cancellation reason.
type CancellationReason17Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason14Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CancellationReason17Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason14Code)(&value)
}

func (c *CancellationReason17Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
