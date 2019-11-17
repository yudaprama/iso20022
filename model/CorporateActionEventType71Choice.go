package model

// Choice of format for the repair reason.
type CorporateActionEventType71Choice struct {

	// Corporate action event type expressed as an ISO 20022 code.
	Code *CorporateActionEventType24Code `xml:"Cd"`

	// Corporate action event type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType71Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType24Code)(&value)
}

func (c *CorporateActionEventType71Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
