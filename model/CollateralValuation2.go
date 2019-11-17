package model

// Provides details about the valuation of each piece of collateral that is posted.
type CollateralValuation2 struct {

	// Provides the identification of the valued collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Specifies the type of collateral used.
	CollateralType *CollateralType1Code `xml:"CollTp"`

	// Provides the status of settlement of an instruction/financial instrument movement.
	SettlementStatus *SettlementStatus2Code `xml:"SttlmSts"`

	// Specifies the number of days used for interest calculation.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd"`

	// Provides details on the collateral valuation.
	ValuationAmounts *CollateralAmount1 `xml:"ValtnAmts"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethod2Code `xml:"DayCntBsis"`

	// Specifies the exchange rate between the currency of the collateral and the reporting currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies the haircut or valuation factor on the currency of the collateral expressed as a percentage.
	CurrencyHaircut *BaseOneRate `xml:"CcyHrcut,omitempty"`

	// Percentage by which the collateral amount needs to be adjusted.
	AdjustedRate *BaseOneRate `xml:"AdjstdRate,omitempty"`

	// Provides details on the securities collateral.
	SecuritiesCollateral *SecuritiesCollateral2 `xml:"SctiesColl,omitempty"`

	// Provides details on the cash collateral valuation.
	CashCollateral *CashCollateral4 `xml:"CshColl,omitempty"`

	// Provides details on other collateral valuation.
	OtherCollateral *OtherCollateral3 `xml:"OthrColl,omitempty"`
}

func (c *CollateralValuation2) SetCollateralIdentification(value string) {
	c.CollateralIdentification = (*Max35Text)(&value)
}

func (c *CollateralValuation2) SetCollateralType(value string) {
	c.CollateralType = (*CollateralType1Code)(&value)
}

func (c *CollateralValuation2) SetSettlementStatus(value string) {
	c.SettlementStatus = (*SettlementStatus2Code)(&value)
}

func (c *CollateralValuation2) SetNumberOfDaysAccrued(value string) {
	c.NumberOfDaysAccrued = (*Number)(&value)
}

func (c *CollateralValuation2) AddValuationAmounts() *CollateralAmount1 {
	c.ValuationAmounts = new(CollateralAmount1)
	return c.ValuationAmounts
}

func (c *CollateralValuation2) SetDayCountBasis(value string) {
	c.DayCountBasis = (*InterestComputationMethod2Code)(&value)
}

func (c *CollateralValuation2) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CollateralValuation2) SetCurrencyHaircut(value string) {
	c.CurrencyHaircut = (*BaseOneRate)(&value)
}

func (c *CollateralValuation2) SetAdjustedRate(value string) {
	c.AdjustedRate = (*BaseOneRate)(&value)
}

func (c *CollateralValuation2) AddSecuritiesCollateral() *SecuritiesCollateral2 {
	c.SecuritiesCollateral = new(SecuritiesCollateral2)
	return c.SecuritiesCollateral
}

func (c *CollateralValuation2) AddCashCollateral() *CashCollateral4 {
	c.CashCollateral = new(CashCollateral4)
	return c.CashCollateral
}

func (c *CollateralValuation2) AddOtherCollateral() *OtherCollateral3 {
	c.OtherCollateral = new(OtherCollateral3)
	return c.OtherCollateral
}
