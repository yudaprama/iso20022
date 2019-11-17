package model

// Specifies the identification of the collateral account.
type CollateralAccountIdentificationType1Choice struct {

	// Indicates the Type of Collateral Account expressed as a code.
	Type *CollateralAccountType1Code `xml:"Tp,omitempty"`

	// Specifies the collateral account expressed as a proprietary code.
	Proprietary *GenericIdentification29 `xml:"Prtry"`
}

func (c *CollateralAccountIdentificationType1Choice) SetType(value string) {
	c.Type = (*CollateralAccountType1Code)(&value)
}

func (c *CollateralAccountIdentificationType1Choice) AddProprietary() *GenericIdentification29 {
	c.Proprietary = new(GenericIdentification29)
	return c.Proprietary
}
