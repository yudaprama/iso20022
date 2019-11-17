package model

// Choice between a standard code or a proprietary code to specify the narrative type of corporate action.
type CorporateActionNarrative4Choice struct {

	// Standard code to specify the narrative type of the message.
	Code *CorporateActionNarrative1Code `xml:"Cd"`

	// Proprietary identification of the narrative type of the message.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionNarrative4Choice) SetCode(value string) {
	c.Code = (*CorporateActionNarrative1Code)(&value)
}

func (c *CorporateActionNarrative4Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
