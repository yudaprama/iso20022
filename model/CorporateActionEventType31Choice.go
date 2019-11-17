package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType31Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType16Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType31Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType16Code)(&value)
}

func (c *CorporateActionEventType31Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
