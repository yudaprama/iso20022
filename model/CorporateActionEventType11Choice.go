package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType11Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType10Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType11Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType10Code)(&value)
}

func (c *CorporateActionEventType11Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
