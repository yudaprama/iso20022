package model

// Choice between a standard code or a proprietary code for specifying the processing type of a corporate action event.
type CorporateActionEventProcessingType2Choice struct {

	// Standard code to specify the processing type of a corporate action event.
	Code *CorporateActionEventProcessingType1Code `xml:"Cd"`

	// Proprietary identification of the processing type of a corporate action event.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventProcessingType2Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventProcessingType1Code)(&value)
}

func (c *CorporateActionEventProcessingType2Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
