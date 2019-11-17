package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType52Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType20Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType52Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType20Code)(&value)
}

func (c *CorporateActionEventType52Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
