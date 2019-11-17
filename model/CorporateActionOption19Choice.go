package model

// Choice between a standard code or a proprietary code.
type CorporateActionOption19Choice struct {

	// Option type expressed as a code.
	Code *CorporateActionOption8Code `xml:"Cd"`

	// Option type expressed as a proprietary identification.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionOption19Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption8Code)(&value)
}

func (c *CorporateActionOption19Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
