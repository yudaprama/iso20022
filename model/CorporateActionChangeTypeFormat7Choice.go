package model

// Choice between a standard code or a proprietary code to specify a type of change.
type CorporateActionChangeTypeFormat7Choice struct {

	// Standard code to specify the type of changes.
	Code *CorporateActionChangeType2Code `xml:"Cd"`

	// Proprietary identification of the type of changes.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionChangeTypeFormat7Choice) SetCode(value string) {
	c.Code = (*CorporateActionChangeType2Code)(&value)
}

func (c *CorporateActionChangeTypeFormat7Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
