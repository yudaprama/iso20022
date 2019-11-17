package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType33Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType17Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType33Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType17Code)(&value)
}

func (c *CorporateActionEventType33Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
