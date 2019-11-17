package model

// Provides additional information on the collateral account of the party delivering/receiving the collateral.
type CollateralAccount2 struct {

	// Unique identification of the collateral account.
	Identification *Max35Text `xml:"Id"`

	// Indicates the type of collateral account.
	Type *CollateralAccountIdentificationType2Choice `xml:"Tp,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CollateralAccount2) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CollateralAccount2) AddType() *CollateralAccountIdentificationType2Choice {
	c.Type = new(CollateralAccountIdentificationType2Choice)
	return c.Type
}

func (c *CollateralAccount2) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}
