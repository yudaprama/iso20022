package model

// Choice between a standard code or a proprietary code to specify the narrative type of corporate action.
type CorporateActionNarrative3Choice struct {

	// Standard code to specify the narrative type of the message.
	Code *CorporateActionNarrative1Code `xml:"Cd"`

	// Proprietary identification of the narrative type of the message.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionNarrative3Choice) SetCode(value string) {
	c.Code = (*CorporateActionNarrative1Code)(&value)
}

func (c *CorporateActionNarrative3Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
