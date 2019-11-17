package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType12Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType11Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType12Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType11Code)(&value)
}

func (c *CorporateActionEventType12Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
