package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType58Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType20Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType58Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType20Code)(&value)
}

func (c *CorporateActionEventType58Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
