package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType46Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType19Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType46Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType19Code)(&value)
}

func (c *CorporateActionEventType46Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
