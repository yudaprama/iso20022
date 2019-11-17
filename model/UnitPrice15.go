package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice15 struct {

	// Type and information about a price.
	Type *TypeOfPrice9Code `xml:"Tp"`

	// Type and information about a price.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

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

	// Indicates whether the price is an estimated price.
	EstimatedPriceIndicator *YesNoIndicator `xml:"EstmtdPricInd"`

	// Specifies the number of days from trade date that the counterparty on the other side of the trade should "given up" or divulged.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerShr,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated2Code `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	ExtendedTaxableIncomePerShareCalculated *Extended350Code `xml:"XtndedTaxblIncmPerShrClctd,omitempty"`

	// Amount included in the dividend that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerDividend *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerDvdd,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUDividendStatus *EUDividendStatus1Code `xml:"EUDvddSts,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	ExtendedEUDividendStatus *Extended350Code `xml:"XtndedEUDvddSts,omitempty"`

	// Amount of money associated with a service.
	ChargeDetails []*Charge15 `xml:"ChrgDtls,omitempty"`

	// Information related to taxes that are due.
	TaxLiabilityDetails []*Tax17 `xml:"TaxLbltyDtls,omitempty"`

	// Information related to taxes that are paid back.
	TaxRefundDetails []*Tax17 `xml:"TaxRfndDtls,omitempty"`
}

func (u *UnitPrice15) SetType(value string) {
	u.Type = (*TypeOfPrice9Code)(&value)
}

func (u *UnitPrice15) SetExtendedType(value string) {
	u.ExtendedType = (*Extended350Code)(&value)
}

func (u *UnitPrice15) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice15) AddValueInInvestmentCurrency() *PriceValue1 {
	newValue := new(PriceValue1)
	u.ValueInInvestmentCurrency = append(u.ValueInInvestmentCurrency, newValue)
	return newValue
}

func (u *UnitPrice15) AddValueInAlternativeCurrency() *PriceValue1 {
	newValue := new(PriceValue1)
	u.ValueInAlternativeCurrency = append(u.ValueInAlternativeCurrency, newValue)
	return newValue
}

func (u *UnitPrice15) SetForExecutionIndicator(value string) {
	u.ForExecutionIndicator = (*YesNoIndicator)(&value)
}

func (u *UnitPrice15) SetCumDividendIndicator(value string) {
	u.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (u *UnitPrice15) SetCalculationBasis(value string) {
	u.CalculationBasis = (*PercentageRate)(&value)
}

func (u *UnitPrice15) SetEstimatedPriceIndicator(value string) {
	u.EstimatedPriceIndicator = (*YesNoIndicator)(&value)
}

func (u *UnitPrice15) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice15) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (u *UnitPrice15) SetTaxableIncomePerShareCalculated(value string) {
	u.TaxableIncomePerShareCalculated = (*TaxableIncomePerShareCalculated2Code)(&value)
}

func (u *UnitPrice15) SetExtendedTaxableIncomePerShareCalculated(value string) {
	u.ExtendedTaxableIncomePerShareCalculated = (*Extended350Code)(&value)
}

func (u *UnitPrice15) SetTaxableIncomePerDividend(value, currency string) {
	u.TaxableIncomePerDividend = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (u *UnitPrice15) SetEUDividendStatus(value string) {
	u.EUDividendStatus = (*EUDividendStatus1Code)(&value)
}

func (u *UnitPrice15) SetExtendedEUDividendStatus(value string) {
	u.ExtendedEUDividendStatus = (*Extended350Code)(&value)
}

func (u *UnitPrice15) AddChargeDetails() *Charge15 {
	newValue := new(Charge15)
	u.ChargeDetails = append(u.ChargeDetails, newValue)
	return newValue
}

func (u *UnitPrice15) AddTaxLiabilityDetails() *Tax17 {
	newValue := new(Tax17)
	u.TaxLiabilityDetails = append(u.TaxLiabilityDetails, newValue)
	return newValue
}

func (u *UnitPrice15) AddTaxRefundDetails() *Tax17 {
	newValue := new(Tax17)
	u.TaxRefundDetails = append(u.TaxRefundDetails, newValue)
	return newValue
}
