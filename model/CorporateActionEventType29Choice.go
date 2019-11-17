package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType29Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType19Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType29Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType19Code)(&value)
}

func (c *CorporateActionEventType29Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
