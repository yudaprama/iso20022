package model

// Choice between a standard code or proprietary code to specify the consent type format required.
type ConsentTypeFormat1Choice struct {

	// Standard code to specify the consent type required.
	Code *ConsentType1Code `xml:"Cd"`

	// Proprietary identification of the consent type.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *ConsentTypeFormat1Choice) SetCode(value string) {
	c.Code = (*ConsentType1Code)(&value)
}

func (c *ConsentTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
