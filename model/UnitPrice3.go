package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice3 struct {

	// Type and information about a price.
	PriceType *TypeOfPrice2Code `xml:"PricTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Interest that has accumulated between the most recent payment of interest and the sale of the financial instrument.
	AccruedInterestNAV *CurrencyAndAmount `xml:"AcrdIntrstNAV,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *CurrencyAndAmount `xml:"TaxblIncmPerShr,omitempty"`
}

func (u *UnitPrice3) SetPriceType(value string) {
	u.PriceType = (*TypeOfPrice2Code)(&value)
}

func (u *UnitPrice3) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice3) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice3) SetAccruedInterestNAV(value, currency string) {
	u.AccruedInterestNAV = NewCurrencyAndAmount(value, currency)
}

func (u *UnitPrice3) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice3) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewCurrencyAndAmount(value, currency)
}
