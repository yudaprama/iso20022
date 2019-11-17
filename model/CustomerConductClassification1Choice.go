package model

// Choice of formats for the specification of the customer conduct classification.
type CustomerConductClassification1Choice struct {

	// Conduct type expressed as a code.
	Code *ConductClassification1Code `xml:"Cd"`

	// Conduct type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CustomerConductClassification1Choice) SetCode(value string) {
	c.Code = (*ConductClassification1Code)(&value)
}

func (c *CustomerConductClassification1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
