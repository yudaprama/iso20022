package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType3Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType6Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType3Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType6Code)(&value)
}

func (c *CorporateActionEventType3Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
