package model

// Choice of formats to  express the stage of the CA event.
type CorporateActionEventStage1FormatChoice struct {

	// Standard code to  specify the stage of the CA event.
	Code *CorporateActionEventStage1Code `xml:"Cd"`

	// Proprietary code to  express the stage of the CA event.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CorporateActionEventStage1FormatChoice) SetCode(value string) {
	c.Code = (*CorporateActionEventStage1Code)(&value)
}

func (c *CorporateActionEventStage1FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
