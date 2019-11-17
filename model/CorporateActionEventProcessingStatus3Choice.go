package model

// Choice of corporate action processing status.
type CorporateActionEventProcessingStatus3Choice struct {

	// Status of a corporate action or the status of a payment expressed as an ISO 20022 code.
	Code *CorporateActionEventProcessingStatus1Code `xml:"Cd"`

	// Status of a corporate action or the status of a payment expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventProcessingStatus3Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventProcessingStatus1Code)(&value)
}

func (c *CorporateActionEventProcessingStatus3Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
