package model

// Choice between a standard code or a proprietary code to specify the type of corporate action options.
type CorporateActionOption23Choice struct {

	// Standard code to specify the type of corporate action options.
	Code *CorporateActionOption7Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action options.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionOption23Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption7Code)(&value)
}

func (c *CorporateActionOption23Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
