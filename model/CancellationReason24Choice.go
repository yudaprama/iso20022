package model

// Choice of format for the cancellation reason.
type CancellationReason24Choice struct {

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Code *CancelledStatusReason13Code `xml:"Cd"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CancellationReason24Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason13Code)(&value)
}

func (c *CancellationReason24Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
