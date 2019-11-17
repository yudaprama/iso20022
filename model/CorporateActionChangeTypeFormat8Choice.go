package model

// Choice between a standard code or proprietary code to specify the change type of a corporate action event.
type CorporateActionChangeTypeFormat8Choice struct {

	// Standard code to specify the type of changes.
	Code *CorporateActionChangeType1Code `xml:"Cd"`

	// Proprietary identification of the type of changes.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionChangeTypeFormat8Choice) SetCode(value string) {
	c.Code = (*CorporateActionChangeType1Code)(&value)
}

func (c *CorporateActionChangeTypeFormat8Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
