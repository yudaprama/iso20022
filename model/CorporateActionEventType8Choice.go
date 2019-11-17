package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType8Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType9Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType8Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType9Code)(&value)
}

func (c *CorporateActionEventType8Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
