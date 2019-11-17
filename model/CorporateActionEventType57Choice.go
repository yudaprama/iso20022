package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType57Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType22Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType57Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType22Code)(&value)
}

func (c *CorporateActionEventType57Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
