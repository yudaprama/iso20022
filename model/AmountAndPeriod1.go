package model

// Relates an amount to a period of time.
type AmountAndPeriod1 struct {

	// Amount of this period.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Start of period or immediate if not specified.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// End of period or indefinite if not specified.
	EndDate *ISODate `xml:"EndDt,omitempty"`
}

func (a *AmountAndPeriod1) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndPeriod1) SetStartDate(value string) {
	a.StartDate = (*ISODate)(&value)
}

func (a *AmountAndPeriod1) SetEndDate(value string) {
	a.EndDate = (*ISODate)(&value)
}
