package model

// Provides the current and market value of the collateral held.
type Collateral6 struct {

	// Value of the collateral after deduction of a percentage (the haircut) that reflects the perceived risk associated with holding this collateral.
	PostHaircutValue *ActiveCurrencyAndAmount `xml:"PstHrcutVal"`

	// Value of the underlying collateral (cash, securities, Letter of credit..) based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal"`

	// Provides the type of collateral, such as securities or cash.
	CollateralType *CollateralType1Code `xml:"CollTp"`
}

func (c *Collateral6) SetPostHaircutValue(value, currency string) {
	c.PostHaircutValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Collateral6) SetMarketValue(value, currency string) {
	c.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Collateral6) SetCollateralType(value string) {
	c.CollateralType = (*CollateralType1Code)(&value)
}
