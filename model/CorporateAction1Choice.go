package model

// Choice of formats for the corporate action event type.
type CorporateAction1Choice struct {

	// Corporate action event  type expressed as a code.
	Type *CorporateActionEventType1Code `xml:"Tp"`

	// Corporate action event expressed as a proprietary code
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateAction1Choice) SetType(value string) {
	c.Type = (*CorporateActionEventType1Code)(&value)
}

func (c *CorporateAction1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
