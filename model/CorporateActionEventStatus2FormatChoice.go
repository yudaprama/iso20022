package model

// Choice of formats to  express the status of the CA option.
type CorporateActionEventStatus2FormatChoice struct {

	// Standard code to specify the status of the CA option.
	Code *CorporateActionEventStatus2Code `xml:"Cd"`

	// Proprietary code to  express the status of the CA option.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CorporateActionEventStatus2FormatChoice) SetCode(value string) {
	c.Code = (*CorporateActionEventStatus2Code)(&value)
}

func (c *CorporateActionEventStatus2FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
