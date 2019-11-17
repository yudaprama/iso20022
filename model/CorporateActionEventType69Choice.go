package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType69Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType24Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType69Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType24Code)(&value)
}

func (c *CorporateActionEventType69Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
