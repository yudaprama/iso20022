package model

// Choice between a standard code or proprietary code to specify the event stage type.
type CorporateActionEventStageFormat20Choice struct {

	// Standard code to specify the stage of the corporate action event.
	Code *CorporateActionEventStage3Code `xml:"Cd"`

	// Proprietary identification of the stage of the corporate action event.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionEventStageFormat20Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventStage3Code)(&value)
}

func (c *CorporateActionEventStageFormat20Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
