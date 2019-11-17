package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice5 struct {

	// Type and information about a price.
	Type *PriceType1 `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *AmountPrice1Choice `xml:"TaxblIncmPerShr,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated1 `xml:"TaxblIncmPerShrClctd,omitempty"`
}

func (u *UnitPrice5) AddType() *PriceType1 {
	u.Type = new(PriceType1)
	return u.Type
}

func (u *UnitPrice5) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice5) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice5) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice5) AddTaxableIncomePerShare() *AmountPrice1Choice {
	u.TaxableIncomePerShare = new(AmountPrice1Choice)
	return u.TaxableIncomePerShare
}

func (u *UnitPrice5) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculated1 {
	u.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculated1)
	return u.TaxableIncomePerShareCalculated
}
