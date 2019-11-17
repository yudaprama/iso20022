package model

// Specifies the status of the details of the corporate action event.
type CorporateActionProcessingStatus4Choice struct {

	// Specifies the status of the details of the corporate action event.
	Code *CorporateActionEventStatus1 `xml:"Cd"`

	// Information related to an identification, for example, party identification or account identification.
	Proprietary *GenericIdentification25 `xml:"Prtry"`
}

func (c *CorporateActionProcessingStatus4Choice) AddCode() *CorporateActionEventStatus1 {
	c.Code = new(CorporateActionEventStatus1)
	return c.Code
}

func (c *CorporateActionProcessingStatus4Choice) AddProprietary() *GenericIdentification25 {
	c.Proprietary = new(GenericIdentification25)
	return c.Proprietary
}
