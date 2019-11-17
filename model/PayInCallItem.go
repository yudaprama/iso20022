package model

// Numerical representation of the net increases and decreases in an account at a specific point in time. A cash balance is calculated from a sum of cash credits minus a sum of cash debits.
type PayInCallItem struct {

	// Specifies the currency and amount of the called item.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (p *PayInCallItem) SetAmount(value, currency string) {
	p.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
