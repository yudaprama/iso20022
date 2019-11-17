package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType54Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType23Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType54Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType23Code)(&value)
}

func (c *CorporateActionEventType54Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
