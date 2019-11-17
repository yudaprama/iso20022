package model

// Choice of format for the repair reason.
type CorporateActionEventType56Choice struct {

	// Corporate action event type expressed as an ISO 20022 code.
	Code *CorporateActionEventType24Code `xml:"Cd"`

	// Corporate action event type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventType56Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType24Code)(&value)
}

func (c *CorporateActionEventType56Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
