package model

// Choice between a standard code or a proprietary code.
type CorporateActionOption11Choice struct {

	// Option type expressed as a code.
	Code *CorporateActionOption8Code `xml:"Cd"`

	// Option type expressed as a proprietary identification.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionOption11Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption8Code)(&value)
}

func (c *CorporateActionOption11Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
