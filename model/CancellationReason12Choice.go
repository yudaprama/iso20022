package model

// Choice of format for the cancellation reason.
type CancellationReason12Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason13Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CancellationReason12Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason13Code)(&value)
}

func (c *CancellationReason12Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
