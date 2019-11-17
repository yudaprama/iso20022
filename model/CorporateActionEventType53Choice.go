package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType53Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType21Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType53Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType21Code)(&value)
}

func (c *CorporateActionEventType53Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
