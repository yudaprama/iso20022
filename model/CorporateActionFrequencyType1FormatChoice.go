package model

// Choice of formats to  express the frequency of a CA event.
type CorporateActionFrequencyType1FormatChoice struct {

	// Standard code to  specify the frequency of a CA event.
	Code *CorporateActionFrequencyType1Code `xml:"Cd"`

	// Proprietary code to  express the frequency of a CA event.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CorporateActionFrequencyType1FormatChoice) SetCode(value string) {
	c.Code = (*CorporateActionFrequencyType1Code)(&value)
}

func (c *CorporateActionFrequencyType1FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
