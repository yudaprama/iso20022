package model

// Choice between a standard code or proprietary code to specify the event stage type.
type CorporateActionEventStageFormat13Choice struct {

	// Standard code to specify the stage of the corporate action event.
	Code *CorporateActionEventStage3Code `xml:"Cd"`

	// Proprietary identification of the stage of the corporate action event.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionEventStageFormat13Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventStage3Code)(&value)
}

func (c *CorporateActionEventStageFormat13Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
