package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice12 struct {

	// Type and information about a price.
	Type *TypeOfPrice12Code `xml:"Tp"`

	// Type and information about a price.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Value of the price, eg, as a currency and value.
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

func (u *UnitPrice12) SetType(value string) {
	u.Type = (*TypeOfPrice12Code)(&value)
}

func (u *UnitPrice12) SetExtendedType(value string) {
	u.ExtendedType = (*Extended350Code)(&value)
}

func (u *UnitPrice12) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice12) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice12) SetAccruedInterestNAV(value, currency string) {
	u.AccruedInterestNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (u *UnitPrice12) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice12) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
