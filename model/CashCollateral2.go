package model

// Provides details about the cash posted as collateral.
type CashCollateral2 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Provides the identification of the clearing member 's cash account.
	CashAccountIdentification *AccountIdentification4Choice `xml:"CshAcctId,omitempty"`

	// Indicates whether excess of cash should be returned or not.
	ReturnExcess *YesNoIndicator `xml:"RtrXcss,omitempty"`

	// Amount of the deposit.
	DepositAmount *ActiveCurrencyAndAmount `xml:"DpstAmt,omitempty"`

	// Specifies whether the deposit is fixed term or call/notice.
	DepositType *DepositType1Code `xml:"DpstTp,omitempty"`

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

func (c *CashCollateral2) SetCollateralIdentification(value string) {
	c.CollateralIdentification = (*Max35Text)(&value)
}

func (c *CashCollateral2) AddCashAccountIdentification() *AccountIdentification4Choice {
	c.CashAccountIdentification = new(AccountIdentification4Choice)
	return c.CashAccountIdentification
}

func (c *CashCollateral2) SetReturnExcess(value string) {
	c.ReturnExcess = (*YesNoIndicator)(&value)
}

func (c *CashCollateral2) SetDepositAmount(value, currency string) {
	c.DepositAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral2) SetDepositType(value string) {
	c.DepositType = (*DepositType1Code)(&value)
}

func (c *CashCollateral2) SetMaturityDate(value string) {
	c.MaturityDate = (*ISODate)(&value)
}

func (c *CashCollateral2) SetValueDate(value string) {
	c.ValueDate = (*ISODate)(&value)
}

func (c *CashCollateral2) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CashCollateral2) SetCollateralValue(value, currency string) {
	c.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral2) SetHaircut(value string) {
	c.Haircut = (*PercentageRate)(&value)
}
