package model

// Choice of format for the repair reason.
type CorporateActionEventType1Choice struct {

	// Corporate action event type expressed as an ISO 20022 code.
	Code *CorporateActionEventType3Code `xml:"Cd"`

	// Corporate action event type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType1Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType3Code)(&value)
}

func (c *CorporateActionEventType1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
