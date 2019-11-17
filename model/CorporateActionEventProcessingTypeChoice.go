package model

// Choice between a standard code or a proprietary code for specifying the processing type of a corporate action event.
type CorporateActionEventProcessingTypeChoice struct {

	// Standard code to specify the processing type of a corporate action event.
	Code *CorporateActionEventProcessingType1Code `xml:"Cd"`

	// Proprietary identification of the processing type of a corporate action event.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventProcessingTypeChoice) SetCode(value string) {
	c.Code = (*CorporateActionEventProcessingType1Code)(&value)
}

func (c *CorporateActionEventProcessingTypeChoice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
