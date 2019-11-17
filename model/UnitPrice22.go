package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice22 struct {

	// Type and information about a price.
	Type *TypeOfPrice46Choice `xml:"Tp"`

	// Value of the price, for example, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerShr,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated2Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Reason why the price is different from the current market price.
	PriceDifferenceReason *Max350Text `xml:"PricDiffRsn,omitempty"`
}

func (u *UnitPrice22) AddType() *TypeOfPrice46Choice {
	u.Type = new(TypeOfPrice46Choice)
	return u.Type
}

func (u *UnitPrice22) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice22) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice22) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice22) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (u *UnitPrice22) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculated2Choice {
	u.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculated2Choice)
	return u.TaxableIncomePerShareCalculated
}

func (u *UnitPrice22) SetPriceDifferenceReason(value string) {
	u.PriceDifferenceReason = (*Max350Text)(&value)
}
