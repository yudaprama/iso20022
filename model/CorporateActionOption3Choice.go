package model

// Choice between a standard code or a proprietary code.
type CorporateActionOption3Choice struct {

	// Option type expressed as a code.
	Code *CorporateActionOption3Code `xml:"Cd"`

	// Option type expressed as a proprietary identification.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionOption3Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption3Code)(&value)
}

func (c *CorporateActionOption3Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
