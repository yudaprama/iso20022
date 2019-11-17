package model

// Choice between a standard code or a proprietary code to specify the type of corporate action options.
type CorporateActionOption10Choice struct {

	// Standard code to specify the type of corporate action options.
	Code *CorporateActionOption7Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action options.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionOption10Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption7Code)(&value)
}

func (c *CorporateActionOption10Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
