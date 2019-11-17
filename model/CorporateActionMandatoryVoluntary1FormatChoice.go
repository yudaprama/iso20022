package model

// Choice of formats to  express whether the event is mandatory, mandatory with options or voluntary.
type CorporateActionMandatoryVoluntary1FormatChoice struct {

	// Standard code to  specify whether the event is mandatory, mandatory with options or voluntary.
	Code *CorporateActionMandatoryVoluntary1Code `xml:"Cd"`

	// Proprietary code to  express whether the event is mandatory, mandatory with options or voluntary.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CorporateActionMandatoryVoluntary1FormatChoice) SetCode(value string) {
	c.Code = (*CorporateActionMandatoryVoluntary1Code)(&value)
}

func (c *CorporateActionMandatoryVoluntary1FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
