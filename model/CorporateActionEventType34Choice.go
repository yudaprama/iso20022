package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType34Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType18Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType34Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType18Code)(&value)
}

func (c *CorporateActionEventType34Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
