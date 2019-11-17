package model

// Choice between a standard code or proprietary code to specify the event stage type.
type CorporateActionEventStageFormat14Choice struct {

	// Standard code to specify the stage of the corporate action event.
	Code *CorporateActionEventStage4Code `xml:"Cd"`

	// Proprietary identification of the stage of the corporate action event.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventStageFormat14Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventStage4Code)(&value)
}

func (c *CorporateActionEventStageFormat14Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
