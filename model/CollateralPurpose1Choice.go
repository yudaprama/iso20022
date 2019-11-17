package model

// Choice between a code and a proprietary code for collateral purpose.
type CollateralPurpose1Choice struct {

	// Provides the collateral purpose using an ISO 20022 code.
	Code *CollateralPurpose1Code `xml:"Cd"`

	// Provides the collateral purpose using a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CollateralPurpose1Choice) SetCode(value string) {
	c.Code = (*CollateralPurpose1Code)(&value)
}

func (c *CollateralPurpose1Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
