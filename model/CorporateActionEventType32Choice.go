package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType32Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType15Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType32Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType15Code)(&value)
}

func (c *CorporateActionEventType32Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
