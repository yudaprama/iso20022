package model

// Withdrawal limits for the account.
type ATMTransactionAmounts4 struct {

	// True if limits may be displayed on the ATM to the customer.
	DisplayFlag *TrueFalseIndicator `xml:"DispFlg,omitempty"`

	// Amount available for withdrawal on the account.
	AvailableAmount *ImpliedCurrencyAndAmount `xml:"AvlblAmt,omitempty"`

	// Remaining daily amount of the customer totals for withdrawals on the account.
	DailyBalance *DetailedAmount4 `xml:"DalyBal,omitempty"`

	// Remaining weekly amount of the customer totals for withdrawals on the account.
	WeeklyBalance *DetailedAmount4 `xml:"WklyBal,omitempty"`

	// Remaining monthly amount of the customer totals for withdrawals on the account.
	MonthlyBalance *DetailedAmount4 `xml:"MnthlyBal,omitempty"`
}

func (a *ATMTransactionAmounts4) SetDisplayFlag(value string) {
	a.DisplayFlag = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransactionAmounts4) SetAvailableAmount(value, currency string) {
	a.AvailableAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts4) AddDailyBalance() *DetailedAmount4 {
	a.DailyBalance = new(DetailedAmount4)
	return a.DailyBalance
}

func (a *ATMTransactionAmounts4) AddWeeklyBalance() *DetailedAmount4 {
	a.WeeklyBalance = new(DetailedAmount4)
	return a.WeeklyBalance
}

func (a *ATMTransactionAmounts4) AddMonthlyBalance() *DetailedAmount4 {
	a.MonthlyBalance = new(DetailedAmount4)
	return a.MonthlyBalance
}
