package model

// Choice of format between an ISO 20022 code and a proprietary code.
type CodeOrProprietary1Choice struct {

	// Element expressed as an ISO 20022 code from an external list.
	Code *Max4Text `xml:"Cd"`

	// Element expressed as a proprietary code.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CodeOrProprietary1Choice) SetCode(value string) {
	c.Code = (*Max4Text)(&value)
}

func (c *CodeOrProprietary1Choice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
