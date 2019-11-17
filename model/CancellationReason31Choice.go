package model

// Choice of format for the cancellation reason.
type CancellationReason31Choice struct {

	// Confirmation cancellation reason expressed as a code.
	Code *ConfirmationCancellationReason1Code `xml:"Cd"`

	// Confirmation cancellation reason expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CancellationReason31Choice) SetCode(value string) {
	c.Code = (*ConfirmationCancellationReason1Code)(&value)
}

func (c *CancellationReason31Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
