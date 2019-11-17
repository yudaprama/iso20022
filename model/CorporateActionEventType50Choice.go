package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType50Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType15Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType50Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType15Code)(&value)
}

func (c *CorporateActionEventType50Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
