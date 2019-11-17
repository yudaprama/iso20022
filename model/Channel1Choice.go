package model

// Choice of format for the submission channel.
type Channel1Choice struct {

	// Submission channel.
	//
	Code *ExternalChannel1Code `xml:"Cd"`

	// Submission channel expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (c *Channel1Choice) SetCode(value string) {
	c.Code = (*ExternalChannel1Code)(&value)
}

func (c *Channel1Choice) AddProprietary() *GenericIdentification1 {
	c.Proprietary = new(GenericIdentification1)
	return c.Proprietary
}
