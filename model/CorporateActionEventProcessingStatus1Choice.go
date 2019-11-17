package model

// Choice of corporate action processing status.
type CorporateActionEventProcessingStatus1Choice struct {

	// Status of a corporate action or the status of a payment expressed as an ISO 20022 code.
	Code *CorporateActionEventProcessingStatus1Code `xml:"Cd"`

	// Status of a corporate action or the status of a payment expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventProcessingStatus1Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventProcessingStatus1Code)(&value)
}

func (c *CorporateActionEventProcessingStatus1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
