package model

// Choice between a standard code or proprietary code to specify the type of conversion.
type ConversionTypeFormat3Choice struct {

	// Standard code to specify the type of conversion.
	Code *ConversionType1Code `xml:"Cd"`

	// Proprietary identification of the type of conversion.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *ConversionTypeFormat3Choice) SetCode(value string) {
	c.Code = (*ConversionType1Code)(&value)
}

func (c *ConversionTypeFormat3Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
