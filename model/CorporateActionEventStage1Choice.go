package model

// Choice of corporate action stage.
type CorporateActionEventStage1Choice struct {

	// Corporate action event stage expressed as an ISO 20022 code.
	Code *CorporateActionEventStage2Code `xml:"Cd"`

	// Corporate action event stage expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventStage1Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventStage2Code)(&value)
}

func (c *CorporateActionEventStage1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
