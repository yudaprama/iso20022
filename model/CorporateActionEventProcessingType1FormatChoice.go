package model

// Choice of formats to  express the type of event processing.
type CorporateActionEventProcessingType1FormatChoice struct {

	// Standard code to  specify the type of event processing.
	Code *CorporateActionEventProcessingType1Code `xml:"Cd"`

	// Proprietary code to  express the type of event processing.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CorporateActionEventProcessingType1FormatChoice) SetCode(value string) {
	c.Code = (*CorporateActionEventProcessingType1Code)(&value)
}

func (c *CorporateActionEventProcessingType1FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
