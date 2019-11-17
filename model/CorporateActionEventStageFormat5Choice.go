package model

// Choice between a standard code or proprietary code to specify the event stage type.
type CorporateActionEventStageFormat5Choice struct {

	// Standard code to specify the stage of the corporate action event.
	Code *CorporateActionEventStage3Code `xml:"Cd"`

	// Proprietary identification of the stage of the corporate action event.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventStageFormat5Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventStage3Code)(&value)
}

func (c *CorporateActionEventStageFormat5Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
