package model

// Choice of format for the repair reason.
type CorporateActionEventType30Choice struct {

	// Corporate action event type expressed as an ISO 20022 code.
	Code *CorporateActionEventType19Code `xml:"Cd"`

	// Corporate action event type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType30Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType19Code)(&value)
}

func (c *CorporateActionEventType30Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
