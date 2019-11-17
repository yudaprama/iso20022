package model

// Choice between a standard code or a proprietary code to specify a type of change.
type CorporateActionChangeTypeFormat2Choice struct {

	// Standard code to specify the type of changes.
	Code *CorporateActionChangeType2Code `xml:"Cd"`

	// Proprietary identification of the type of changes.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionChangeTypeFormat2Choice) SetCode(value string) {
	c.Code = (*CorporateActionChangeType2Code)(&value)
}

func (c *CorporateActionChangeTypeFormat2Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
