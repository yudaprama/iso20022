package model

// Provides details about the cash posted as collateral.
type CashCollateral4 struct {

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Amount of the deposit.
	DepositAmount *ActiveCurrencyAndAmount `xml:"DpstAmt,omitempty"`

	// Specifies whether the deposit is fixed term or call/notice.
	DepositType *DepositType1Code `xml:"DpstTp,omitempty"`

	// Amount blocked by the central counterparty for any reasonable reason ( for example for judicial reasons). In this case the investor can not withdraw or distribute this collateral.
	BlockedAmount *ActiveCurrencyAndAmount `xml:"BlckdAmt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Valuation date of the cash taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Exchange rate.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal"`

	// Haircut or valuation factor on the collateral expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`
}

func (c *CashCollateral4) SetAssetNumber(value string) {
	c.AssetNumber = (*Max35Text)(&value)
}

func (c *CashCollateral4) SetDepositAmount(value, currency string) {
	c.DepositAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral4) SetDepositType(value string) {
	c.DepositType = (*DepositType1Code)(&value)
}

func (c *CashCollateral4) SetBlockedAmount(value, currency string) {
	c.BlockedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral4) SetMaturityDate(value string) {
	c.MaturityDate = (*ISODate)(&value)
}

func (c *CashCollateral4) SetValueDate(value string) {
	c.ValueDate = (*ISODate)(&value)
}

func (c *CashCollateral4) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CashCollateral4) SetCollateralValue(value, currency string) {
	c.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral4) SetHaircut(value string) {
	c.Haircut = (*PercentageRate)(&value)
}
