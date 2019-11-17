package model

// Provides information about the collateral valuation such as the collateral amount, the market value.
type CollateralAmount1 struct {

	// Specifies the total amount of the collateral in the collateral currency.
	CollateralAmount *ActiveCurrencyAndAmount `xml:"CollAmt"`

	// Specifies the total amount of the collateral in the reporting currency.
	ReportedCurrencyAndAmount *ActiveCurrencyAndAmount `xml:"RptdCcyAndAmt"`

	// Specifies the total market to market value of the collateral in the reporting currency. It is the dirty price, that is, the accrued interest is included if any.
	MarketValueAmount *ActiveCurrencyAndAmount `xml:"MktValAmt"`

	// Specifies the accrued interest on the value of the collateral in the currency of the collateral.
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the amount of money paid for the provision of financial services.
	FeesAndCommissions *ActiveCurrencyAndAmount `xml:"FeesAndComssns,omitempty"`
}

func (c *CollateralAmount1) SetCollateralAmount(value, currency string) {
	c.CollateralAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralAmount1) SetReportedCurrencyAndAmount(value, currency string) {
	c.ReportedCurrencyAndAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralAmount1) SetMarketValueAmount(value, currency string) {
	c.MarketValueAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralAmount1) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralAmount1) SetFeesAndCommissions(value, currency string) {
	c.FeesAndCommissions = NewActiveCurrencyAndAmount(value, currency)
}
