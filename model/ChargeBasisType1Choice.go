package model

// Choice of formats for the charge basis.
type ChargeBasisType1Choice struct {

	// Charge basis expressed as a code.
	Code *TaxationBasis2Code `xml:"Cd"`

	// Charge basis expressed as a code a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *ChargeBasisType1Choice) SetCode(value string) {
	c.Code = (*TaxationBasis2Code)(&value)
}

func (c *ChargeBasisType1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
