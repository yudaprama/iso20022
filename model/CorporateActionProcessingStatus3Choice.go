package model

// Specifies the status of the details of the corporate action event.
type CorporateActionProcessingStatus3Choice struct {

	// Specifies the status of the details of the corporate action event.
	Code *CorporateActionProcessingStatus1Choice `xml:"Cd"`

	// Information related to an identification, for example, party identification or account identification.
	Proprietary *GenericIdentification25 `xml:"Prtry"`
}

func (c *CorporateActionProcessingStatus3Choice) AddCode() *CorporateActionProcessingStatus1Choice {
	c.Code = new(CorporateActionProcessingStatus1Choice)
	return c.Code
}

func (c *CorporateActionProcessingStatus3Choice) AddProprietary() *GenericIdentification25 {
	c.Proprietary = new(GenericIdentification25)
	return c.Proprietary
}
