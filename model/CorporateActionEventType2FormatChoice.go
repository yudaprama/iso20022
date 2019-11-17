package model

// Choice between formats to express the corporate event type.
type CorporateActionEventType2FormatChoice struct {

	// Event types expressed as a code.
	Code *CorporateActionEventType2Code `xml:"Cd"`

	// Event types expressed as a proprietary code.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CorporateActionEventType2FormatChoice) SetCode(value string) {
	c.Code = (*CorporateActionEventType2Code)(&value)
}

func (c *CorporateActionEventType2FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
