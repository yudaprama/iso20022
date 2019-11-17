package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType5Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType7Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType5Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType7Code)(&value)
}

func (c *CorporateActionEventType5Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
