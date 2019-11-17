package model

// Limit of amounts for the customer.
type ATMTransactionAmounts2 struct {

	// Currency of the limits, if different from the requested amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Maximum amount allowed in the authorised currency if the withdrawal was not approved.
	MaximumAuthorisableAmount *ImpliedCurrencyAndAmount `xml:"MaxAuthsbAmt,omitempty"`

	// Minimum amount allowed for a withdrawal in the authorised currency.
	MinimumAllowedAmount *ImpliedCurrencyAndAmount `xml:"MinAllwdAmt,omitempty"`

	// Maximum amount allowed for a withdrawal in the authorised currency.
	MaximumAllowedAmount *ImpliedCurrencyAndAmount `xml:"MaxAllwdAmt,omitempty"`

	// Remaining daily amount of the customer totals after the withdrawal.
	DailyBalance *DetailedAmount4 `xml:"DalyBal,omitempty"`

	// Remaining weekly amount of the customer totals after the withdrawal.
	WeeklyBalance *DetailedAmount4 `xml:"WklyBal,omitempty"`

	// Remaining monthly amount of the customer totals after the withdrawal.
	MonthlyBalance *DetailedAmount4 `xml:"MnthlyBal,omitempty"`
}

func (a *ATMTransactionAmounts2) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts2) SetMaximumAuthorisableAmount(value, currency string) {
	a.MaximumAuthorisableAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts2) SetMinimumAllowedAmount(value, currency string) {
	a.MinimumAllowedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts2) SetMaximumAllowedAmount(value, currency string) {
	a.MaximumAllowedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts2) AddDailyBalance() *DetailedAmount4 {
	a.DailyBalance = new(DetailedAmount4)
	return a.DailyBalance
}

func (a *ATMTransactionAmounts2) AddWeeklyBalance() *DetailedAmount4 {
	a.WeeklyBalance = new(DetailedAmount4)
	return a.WeeklyBalance
}

func (a *ATMTransactionAmounts2) AddMonthlyBalance() *DetailedAmount4 {
	a.MonthlyBalance = new(DetailedAmount4)
	return a.MonthlyBalance
}
