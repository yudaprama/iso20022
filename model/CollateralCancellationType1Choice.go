package model

// Choice between a code and a proprietary code for the reason the collateral message has been cancelled.
type CollateralCancellationType1Choice struct {

	// Provides the cancellation reason using an ISO 20022 code.
	Code *CollateralManagementCancellationReason1Code `xml:"Cd"`

	// Provides the cancellation reason using a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CollateralCancellationType1Choice) SetCode(value string) {
	c.Code = (*CollateralManagementCancellationReason1Code)(&value)
}

func (c *CollateralCancellationType1Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
