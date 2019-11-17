package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType41Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType16Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType41Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType16Code)(&value)
}

func (c *CorporateActionEventType41Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
