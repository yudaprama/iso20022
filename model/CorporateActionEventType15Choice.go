package model

// Choice of format for the repair reason.
type CorporateActionEventType15Choice struct {

	// Corporate action event type expressed as an ISO 20022 code.
	Code *CorporateActionEventType14Code `xml:"Cd"`

	// Corporate action event type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventType15Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType14Code)(&value)
}

func (c *CorporateActionEventType15Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
