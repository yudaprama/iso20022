package model

// Choice between a standard code or proprietary code to specify the change type of a corporate action event.
type CorporateActionChangeTypeFormat5Choice struct {

	// Standard code to specify the type of changes.
	Code *CorporateActionChangeType1Code `xml:"Cd"`

	// Proprietary identification of the type of changes.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionChangeTypeFormat5Choice) SetCode(value string) {
	c.Code = (*CorporateActionChangeType1Code)(&value)
}

func (c *CorporateActionChangeTypeFormat5Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
