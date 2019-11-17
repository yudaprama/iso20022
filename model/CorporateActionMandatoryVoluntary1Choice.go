package model

// Choice between a standard code or a proprietary code to indicate if a corporate action event is mandatory or not.
type CorporateActionMandatoryVoluntary1Choice struct {

	// Standard code to specify whether the event is mandatory, mandatory with options or voluntary.
	Code *CorporateActionMandatoryVoluntary1Code `xml:"Cd"`

	// Proprietary identification of an event.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionMandatoryVoluntary1Choice) SetCode(value string) {
	c.Code = (*CorporateActionMandatoryVoluntary1Code)(&value)
}

func (c *CorporateActionMandatoryVoluntary1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
