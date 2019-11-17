package model

// Numerical representation of the net increases and decreases in an account at a specific point in time. A cash balance is calculated from a sum of cash credits minus a sum of cash debits.
type BalanceStatus2 struct {

	// Balance in each currency calculated at the value date and time indicated in the report.
	Balance *ActiveCurrencyAndAmount `xml:"Bal"`
}

func (b *BalanceStatus2) SetBalance(value, currency string) {
	b.Balance = NewActiveCurrencyAndAmount(value, currency)
}
