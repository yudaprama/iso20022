package model

// Choice of format for the collateral type.
type CollateralType1Choice struct {

	// Type of collateral expressed as an ISO 20022 code.
	Code *CollateralType3Code `xml:"Cd"`

	// Type of collateral expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (c *CollateralType1Choice) SetCode(value string) {
	c.Code = (*CollateralType3Code)(&value)
}

func (c *CollateralType1Choice) AddProprietary() *GenericIdentification38 {
	c.Proprietary = new(GenericIdentification38)
	return c.Proprietary
}
