package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType7Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType8Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType7Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType8Code)(&value)
}

func (c *CorporateActionEventType7Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
