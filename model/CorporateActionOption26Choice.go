package model

// Choice between a standard code or a proprietary code.
type CorporateActionOption26Choice struct {

	// Option type expressed as a code.
	Code *CorporateActionOption8Code `xml:"Cd"`

	// Option type expressed as a proprietary identification.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionOption26Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption8Code)(&value)
}

func (c *CorporateActionOption26Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
