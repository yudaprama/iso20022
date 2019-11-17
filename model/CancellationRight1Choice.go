package model

// Choice of formats for the cancellation right.
type CancellationRight1Choice struct {

	// Cancellation right expressed as a code.
	Code *CancellationRight1Code `xml:"Cd"`

	// Cancellation right expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CancellationRight1Choice) SetCode(value string) {
	c.Code = (*CancellationRight1Code)(&value)
}

func (c *CancellationRight1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
