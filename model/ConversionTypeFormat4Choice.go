package model

// Choice between a standard code or proprietary code to specify the type of conversion.
type ConversionTypeFormat4Choice struct {

	// Standard code to specify the type of conversion.
	Code *ConversionType1Code `xml:"Cd"`

	// Proprietary identification of the type of conversion.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *ConversionTypeFormat4Choice) SetCode(value string) {
	c.Code = (*ConversionType1Code)(&value)
}

func (c *ConversionTypeFormat4Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
