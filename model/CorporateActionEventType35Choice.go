package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType35Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType17Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType35Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType17Code)(&value)
}

func (c *CorporateActionEventType35Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
