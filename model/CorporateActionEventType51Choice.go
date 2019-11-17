package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType51Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType22Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType51Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType22Code)(&value)
}

func (c *CorporateActionEventType51Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
