package model

// Choice between a standard code or proprietary code to specify the type of conversion.
type ConversionTypeFormat1Choice struct {

	// Standard code to specify the type of conversion.
	Code *ConversionType1Code `xml:"Cd"`

	// Proprietary identification of the type of conversion.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *ConversionTypeFormat1Choice) SetCode(value string) {
	c.Code = (*ConversionType1Code)(&value)
}

func (c *ConversionTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
