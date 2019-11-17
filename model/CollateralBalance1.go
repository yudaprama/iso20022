package model

// Provides details about the collateral held by party A and/or B.
type CollateralBalance1 struct {

	// Collateral currently held by party A.
	HeldByPartyA *ActiveCurrencyAndAmount `xml:"HeldByPtyA"`

	// Collateral currently held by party B.
	HeldByPartyB *ActiveCurrencyAndAmount `xml:"HeldByPtyB"`
}

func (c *CollateralBalance1) SetHeldByPartyA(value, currency string) {
	c.HeldByPartyA = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralBalance1) SetHeldByPartyB(value, currency string) {
	c.HeldByPartyB = NewActiveCurrencyAndAmount(value, currency)
}
