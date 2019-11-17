package model

// Choice of corporate action stage.
type CorporateActionEventStage3Choice struct {

	// Corporate action event stage expressed as an ISO 20022 code.
	Code *CorporateActionEventStage2Code `xml:"Cd"`

	// Corporate action event stage expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventStage3Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventStage2Code)(&value)
}

func (c *CorporateActionEventStage3Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
