package model

// Choice of formats to  express the calculation method for drawings.
type CorporateActionCalculationMethod1FormatChoice struct {

	// Standard code to  specify the calculation method for drawings.
	Code *CorporateActionCalculationMethod1Code `xml:"Cd"`

	// Proprietary code to  express the calculation method for drawings.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CorporateActionCalculationMethod1FormatChoice) SetCode(value string) {
	c.Code = (*CorporateActionCalculationMethod1Code)(&value)
}

func (c *CorporateActionCalculationMethod1FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
