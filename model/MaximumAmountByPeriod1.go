package model

// Maximum amount allowed over a specific period of time.
type MaximumAmountByPeriod1 struct {

	// Maximum amount allowed over a specific period of time.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt"`

	// Period specified as a number of days.
	NumberOfDays *Max3NumericText `xml:"NbOfDays"`
}

func (m *MaximumAmountByPeriod1) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MaximumAmountByPeriod1) SetNumberOfDays(value string) {
	m.NumberOfDays = (*Max3NumericText)(&value)
}
