package model

// Specifies the type of collateral that will be delivered and the date to be reported.
type CollateralMovement9 struct {

	// Specifies the type of collateral.
	CollateralType *CollateralType1Code `xml:"CollTp"`

	// Date by which the collateral movement must be executed.
	Date *ISODate `xml:"Dt,omitempty"`
}

func (c *CollateralMovement9) SetCollateralType(value string) {
	c.CollateralType = (*CollateralType1Code)(&value)
}

func (c *CollateralMovement9) SetDate(value string) {
	c.Date = (*ISODate)(&value)
}
