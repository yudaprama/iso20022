package model

// Choice between a standard code or a proprietary code for specifying the processing type of a corporate action event.
type CorporateActionEventProcessingType3Choice struct {

	// Standard code to specify the processing type of a corporate action event.
	Code *CorporateActionEventProcessingType1Code `xml:"Cd"`

	// Proprietary identification of the processing type of a corporate action event.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventProcessingType3Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventProcessingType1Code)(&value)
}

func (c *CorporateActionEventProcessingType3Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
