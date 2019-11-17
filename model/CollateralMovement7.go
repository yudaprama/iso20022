package model

// Provides the agreed amount and the collateral movement direction.
type CollateralMovement7 struct {

	// Provides the call amount that is agreed and that will result in a collateral movement.
	AgreedAmount *ActiveCurrencyAndAmount `xml:"AgrdAmt"`

	// Provides the collateral movement direction that is a delivery and optionaly a return, or a return only.
	MovementDirection []*CollateralMovement4Choice `xml:"MvmntDrctn,omitempty"`
}

func (c *CollateralMovement7) SetAgreedAmount(value, currency string) {
	c.AgreedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralMovement7) AddMovementDirection() *CollateralMovement4Choice {
	newValue := new(CollateralMovement4Choice)
	c.MovementDirection = append(c.MovementDirection, newValue)
	return newValue
}
