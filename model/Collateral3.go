package model

// Provides the current and market value of the collateral held.
type Collateral3 struct {

	// Value of the collateral after deduction of a percentage (the haircut) that reflects the perceived risk associated with holding this collateral.
	PostHaircutValue *ActiveCurrencyAndAmount `xml:"PstHrcutVal"`

	// Value of the underlying collateral (cash, securities, LoC) based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal"`

	// Provides the type of collateral, such as securities or cash.
	CollateralType *CollateralType2Code `xml:"CollTp"`
}

func (c *Collateral3) SetPostHaircutValue(value, currency string) {
	c.PostHaircutValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Collateral3) SetMarketValue(value, currency string) {
	c.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Collateral3) SetCollateralType(value string) {
	c.CollateralType = (*CollateralType2Code)(&value)
}
