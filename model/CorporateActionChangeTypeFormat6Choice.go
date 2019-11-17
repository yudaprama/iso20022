package model

// Choice between a standard code or a proprietary code to specify a type of change.
type CorporateActionChangeTypeFormat6Choice struct {

	// Standard code to specify the type of changes.
	Code *CorporateActionChangeType2Code `xml:"Cd"`

	// Proprietary identification of the type of changes.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionChangeTypeFormat6Choice) SetCode(value string) {
	c.Code = (*CorporateActionChangeType2Code)(&value)
}

func (c *CorporateActionChangeTypeFormat6Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
