package model

// Choice between a standard code or a proprietary code to indicate if a corporate action event is mandatory or not.
type CorporateActionMandatoryVoluntary3Choice struct {

	// Standard code to specify whether the event is mandatory, mandatory with options or voluntary.
	Code *CorporateActionMandatoryVoluntary1Code `xml:"Cd"`

	// Proprietary identification of an event.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionMandatoryVoluntary3Choice) SetCode(value string) {
	c.Code = (*CorporateActionMandatoryVoluntary1Code)(&value)
}

func (c *CorporateActionMandatoryVoluntary3Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
