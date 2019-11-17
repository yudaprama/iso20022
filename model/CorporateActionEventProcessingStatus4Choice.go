package model

// Choice of corporate action processing status.
type CorporateActionEventProcessingStatus4Choice struct {

	// Status of a corporate action or the status of a payment expressed as an ISO 20022 code.
	Code *CorporateActionEventProcessingStatus1Code `xml:"Cd"`

	// Status of a corporate action or the status of a payment expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventProcessingStatus4Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventProcessingStatus1Code)(&value)
}

func (c *CorporateActionEventProcessingStatus4Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
