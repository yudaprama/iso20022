package model

// Choice between a standard code or proprietary code to specify the consent type format required.
type ConsentTypeFormat4Choice struct {

	// Standard code to specify the consent type required.
	Code *ConsentType1Code `xml:"Cd"`

	// Proprietary identification of the consent type.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *ConsentTypeFormat4Choice) SetCode(value string) {
	c.Code = (*ConsentType1Code)(&value)
}

func (c *ConsentTypeFormat4Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
