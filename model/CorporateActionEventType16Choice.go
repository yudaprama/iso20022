package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType16Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType14Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType16Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType14Code)(&value)
}

func (c *CorporateActionEventType16Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
