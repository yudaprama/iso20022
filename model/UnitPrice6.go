package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice6 struct {

	// Type and information about a price.
	Type *PriceType2 `xml:"Tp"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Value of the price, eg, as a currency and value.
	ValueInInvestmentCurrency []*PriceValue1 `xml:"ValInInvstmtCcy"`

	// Value of the price, eg, as a currency and value.
	ValueInAlternativeCurrency []*PriceValue1 `xml:"ValInAltrntvCcy,omitempty"`

	// Indicates whether the price information can be used for the execution of a transaction.
	ForExecutionIndicator *YesNoIndicator `xml:"ForExctnInd"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Ratio applied on the non-adjusted price.
	CalculationBasis *PercentageRate `xml:"ClctnBsis,omitempty"`

	// Specifies the number of days from trade date that the counterparty on the other side of the trade should "given up" or divulged.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *AmountPrice1Choice `xml:"TaxblIncmPerShr,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated1 `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Amount of money associated with a service.
	ChargeDetails []*Charge9 `xml:"ChrgDtls,omitempty"`

	// Information related to taxes that are due.
	TaxLiabilityDetails []*Tax8 `xml:"TaxLbltyDtls,omitempty"`

	// Information related to taxes that are paid back.
	TaxRefundDetails []*Tax8 `xml:"TaxRfndDtls,omitempty"`
}

func (u *UnitPrice6) AddType() *PriceType2 {
	u.Type = new(PriceType2)
	return u.Type
}

func (u *UnitPrice6) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice6) AddValueInInvestmentCurrency() *PriceValue1 {
	newValue := new(PriceValue1)
	u.ValueInInvestmentCurrency = append(u.ValueInInvestmentCurrency, newValue)
	return newValue
}

func (u *UnitPrice6) AddValueInAlternativeCurrency() *PriceValue1 {
	newValue := new(PriceValue1)
	u.ValueInAlternativeCurrency = append(u.ValueInAlternativeCurrency, newValue)
	return newValue
}

func (u *UnitPrice6) SetForExecutionIndicator(value string) {
	u.ForExecutionIndicator = (*YesNoIndicator)(&value)
}

func (u *UnitPrice6) SetCumDividendIndicator(value string) {
	u.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (u *UnitPrice6) SetCalculationBasis(value string) {
	u.CalculationBasis = (*PercentageRate)(&value)
}

func (u *UnitPrice6) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice6) AddTaxableIncomePerShare() *AmountPrice1Choice {
	u.TaxableIncomePerShare = new(AmountPrice1Choice)
	return u.TaxableIncomePerShare
}

func (u *UnitPrice6) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculated1 {
	u.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculated1)
	return u.TaxableIncomePerShareCalculated
}

func (u *UnitPrice6) AddChargeDetails() *Charge9 {
	newValue := new(Charge9)
	u.ChargeDetails = append(u.ChargeDetails, newValue)
	return newValue
}

func (u *UnitPrice6) AddTaxLiabilityDetails() *Tax8 {
	newValue := new(Tax8)
	u.TaxLiabilityDetails = append(u.TaxLiabilityDetails, newValue)
	return newValue
}

func (u *UnitPrice6) AddTaxRefundDetails() *Tax8 {
	newValue := new(Tax8)
	u.TaxRefundDetails = append(u.TaxRefundDetails, newValue)
	return newValue
}
