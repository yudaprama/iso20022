package model

// Choice between a standard code or a proprietary code to specify the narrative type of corporate action.
type CorporateActionNarrative1Choice struct {

	// Standard code to specify the narrative type of the message.
	Code *CorporateActionNarrative1Code `xml:"Cd"`

	// Proprietary identification of the narrative type of the message.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionNarrative1Choice) SetCode(value string) {
	c.Code = (*CorporateActionNarrative1Code)(&value)
}

func (c *CorporateActionNarrative1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
