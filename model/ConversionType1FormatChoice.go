package model

// Choice of formats to  express the type of conversion.
type ConversionType1FormatChoice struct {

	// Standard code to  specify the type of conversion.
	Code *ConversionType1Code `xml:"Cd"`

	// Proprietary code to  express the type of conversion.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *ConversionType1FormatChoice) SetCode(value string) {
	c.Code = (*ConversionType1Code)(&value)
}

func (c *ConversionType1FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
