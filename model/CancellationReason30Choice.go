package model

// Choice of format for the cancellation reason.
type CancellationReason30Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason12Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CancellationReason30Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason12Code)(&value)
}

func (c *CancellationReason30Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
