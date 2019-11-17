package model

// Choice of corporate action stage.
type CorporateActionEventStage4Choice struct {

	// Corporate action event stage expressed as an ISO 20022 code.
	Code *CorporateActionEventStage2Code `xml:"Cd"`

	// Corporate action event stage expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventStage4Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventStage2Code)(&value)
}

func (c *CorporateActionEventStage4Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
