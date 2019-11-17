package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice21 struct {

	// Type and information about a price.
	Type *TypeOfPrice31Choice `xml:"Tp"`

	// Value of the price, for example, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Interest that has accumulated between the most recent payment of interest and the sale of the financial instrument.
	AccruedInterestNAV *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstNAV,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerShr,omitempty"`
}

func (u *UnitPrice21) AddType() *TypeOfPrice31Choice {
	u.Type = new(TypeOfPrice31Choice)
	return u.Type
}

func (u *UnitPrice21) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice21) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice21) SetAccruedInterestNAV(value, currency string) {
	u.AccruedInterestNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (u *UnitPrice21) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice21) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
