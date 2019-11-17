package model

// Choice between a standard code or proprietary code to specify the change type of a corporate action event.
type CorporateActionChangeTypeFormat1Choice struct {

	// Standard code to specify the type of changes.
	Code *CorporateActionChangeType1Code `xml:"Cd"`

	// Proprietary identification of the type of changes.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionChangeTypeFormat1Choice) SetCode(value string) {
	c.Code = (*CorporateActionChangeType1Code)(&value)
}

func (c *CorporateActionChangeTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
