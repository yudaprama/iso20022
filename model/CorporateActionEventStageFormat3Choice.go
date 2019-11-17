package model

// Choice between a standard code or proprietary code to specify the event stage type.
type CorporateActionEventStageFormat3Choice struct {

	// Standard code to specify the stage of the corporate action event.
	Code *CorporateActionEventStage1Code `xml:"Cd"`

	// Proprietary identification of the stage of the corporate action event.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionEventStageFormat3Choice) SetCode(value string) {
	c.Code = (*CorporateActionEventStage1Code)(&value)
}

func (c *CorporateActionEventStageFormat3Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
