package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType62Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType21Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType62Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType21Code)(&value)
}

func (c *CorporateActionEventType62Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
