package model

// Choice between a standard code or proprietary code to specify the consent type format required.
type ConsentTypeFormat5Choice struct {

	// Standard code to specify the consent type required.
	Code *ConsentType1Code `xml:"Cd"`

	// Proprietary identification of the consent type.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *ConsentTypeFormat5Choice) SetCode(value string) {
	c.Code = (*ConsentType1Code)(&value)
}

func (c *ConsentTypeFormat5Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
