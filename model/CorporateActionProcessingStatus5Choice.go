package model

// Specifies the status of the details of the corporate action event.
type CorporateActionProcessingStatus5Choice struct {

	// Specifies the status of the details of the corporate action event.
	Code *CorporateActionEventStatus1 `xml:"Cd"`

	// Information related to an identification, for example, party identification or account identification.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionProcessingStatus5Choice) AddCode() *CorporateActionEventStatus1 {
	c.Code = new(CorporateActionEventStatus1)
	return c.Code
}

func (c *CorporateActionProcessingStatus5Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
