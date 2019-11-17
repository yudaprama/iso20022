package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType13Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType12Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType13Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType12Code)(&value)
}

func (c *CorporateActionEventType13Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
