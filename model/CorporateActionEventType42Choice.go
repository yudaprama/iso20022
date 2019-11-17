package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType42Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType18Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType42Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType18Code)(&value)
}

func (c *CorporateActionEventType42Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
