package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType14Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType13Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType14Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType13Code)(&value)
}

func (c *CorporateActionEventType14Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
