package model

// Choice between a standard code or a proprietary code to specify the type of corporate action options.
type CorporateActionOption2Choice struct {

	// Standard code to specify the type of corporate action options.
	Code *CorporateActionOption2Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action options.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionOption2Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption2Code)(&value)
}

func (c *CorporateActionOption2Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
