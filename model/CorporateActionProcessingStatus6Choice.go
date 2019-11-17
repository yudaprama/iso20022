package model

// Specifies the status of the details of the corporate action event.
type CorporateActionProcessingStatus6Choice struct {

	// Specifies the status of the details of the corporate action event.
	Code *CorporateActionEventStatus1 `xml:"Cd"`

	// Information related to an identification, for example, party identification or account identification.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionProcessingStatus6Choice) AddCode() *CorporateActionEventStatus1 {
	c.Code = new(CorporateActionEventStatus1)
	return c.Code
}

func (c *CorporateActionProcessingStatus6Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
