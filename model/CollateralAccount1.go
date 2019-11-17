package model

// The Collateral Account provides additional information on the Collateral Account of the Party delivering the collateral.
type CollateralAccount1 struct {

	// Unique identification of the Collateral Account.
	Identification *Max35Text `xml:"Id"`

	// Indicates the Type of Collateral Account.
	Type *CollateralAccountIdentificationType1Choice `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CollateralAccount1) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CollateralAccount1) AddType() *CollateralAccountIdentificationType1Choice {
	c.Type = new(CollateralAccountIdentificationType1Choice)
	return c.Type
}

func (c *CollateralAccount1) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}
