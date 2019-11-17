package model

// Choice between a standard code or a proprietary code to specify the type of corporate action options.
type CorporateActionOption18Choice struct {

	// Standard code to specify the type of corporate action options.
	Code *CorporateActionOption7Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action options.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionOption18Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption7Code)(&value)
}

func (c *CorporateActionOption18Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
