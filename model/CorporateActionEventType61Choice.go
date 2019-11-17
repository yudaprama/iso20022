package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType61Choice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType23Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventType61Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventType23Code)(&value)
}

func (c *CorporateActionEventType61Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
