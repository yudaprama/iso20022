package model

// Specifies the identification of the collateral account.
type CollateralAccountIdentificationType2Choice struct {

	// Indicates the type of collateral account expressed as a code.
	Type *CollateralAccountType1Code `xml:"Tp,omitempty"`

	// Specifies the collateral account expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (c *CollateralAccountIdentificationType2Choice) SetType(value string) {
	c.Type = (*CollateralAccountType1Code)(&value)
}

func (c *CollateralAccountIdentificationType2Choice) AddProprietary() *GenericIdentification36 {
	c.Proprietary = new(GenericIdentification36)
	return c.Proprietary
}
