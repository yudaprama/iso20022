package model

// Specifies the expected collateral type and direction.
type ExpectedCollateralMovement2 struct {

	// Type of collateral that will be delivered and date by which the collateral movement is expected.
	Delivery []*CollateralMovement9 `xml:"Dlvry,omitempty"`

	// Type of collateral that will be returned and date by which the collateral movement is expected.
	Return []*CollateralMovement9 `xml:"Rtr,omitempty"`
}

func (e *ExpectedCollateralMovement2) AddDelivery() *CollateralMovement9 {
	newValue := new(CollateralMovement9)
	e.Delivery = append(e.Delivery, newValue)
	return newValue
}

func (e *ExpectedCollateralMovement2) AddReturn() *CollateralMovement9 {
	newValue := new(CollateralMovement9)
	e.Return = append(e.Return, newValue)
	return newValue
}
