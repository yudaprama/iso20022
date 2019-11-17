package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice10 struct {

	// Type and information about a price.
	Type *TypeOfPrice10Code `xml:"Tp"`

	// Type and information about a price.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerShr,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated2Code `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	ExtendedTaxableIncomePerShareCalculated *Extended350Code `xml:"XtndedTaxblIncmPerShrClctd,omitempty"`

	// Specifies the reason why the price is different from the current market price.
	PriceDifferenceReason *Max350Text `xml:"PricDiffRsn,omitempty"`
}

func (u *UnitPrice10) SetType(value string) {
	u.Type = (*TypeOfPrice10Code)(&value)
}

func (u *UnitPrice10) SetExtendedType(value string) {
	u.ExtendedType = (*Extended350Code)(&value)
}

func (u *UnitPrice10) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice10) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice10) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice10) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (u *UnitPrice10) SetTaxableIncomePerShareCalculated(value string) {
	u.TaxableIncomePerShareCalculated = (*TaxableIncomePerShareCalculated2Code)(&value)
}

func (u *UnitPrice10) SetExtendedTaxableIncomePerShareCalculated(value string) {
	u.ExtendedTaxableIncomePerShareCalculated = (*Extended350Code)(&value)
}

func (u *UnitPrice10) SetPriceDifferenceReason(value string) {
	u.PriceDifferenceReason = (*Max350Text)(&value)
}
