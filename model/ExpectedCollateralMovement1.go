package model

// Specifies the expected collateral type and direction.
type ExpectedCollateralMovement1 struct {

	// Type of collateral that will be delivered.
	Delivery []*CollateralType1Code `xml:"Dlvry,omitempty"`

	// Type of collateral that will be returned.
	Return []*CollateralType1Code `xml:"Rtr,omitempty"`
}

func (e *ExpectedCollateralMovement1) AddDelivery(value string) {
	e.Delivery = append(e.Delivery, (*CollateralType1Code)(&value))
}

func (e *ExpectedCollateralMovement1) AddReturn(value string) {
	e.Return = append(e.Return, (*CollateralType1Code)(&value))
}
