package model

// Aggregated position held in a securities account for a specified financial instrument.
type FinancialInstrumentAggregateBalance2 struct {

	// Balance of settled transactions.
	SettledBalance *FinancialInstrumentQuantity1Choice `xml:"SttldBal,omitempty"`

	// Balance of settled transactions and transactions pending settlement.
	TradedBalance *FinancialInstrumentQuantity1Choice `xml:"TraddBal,omitempty"`

	// Breakdown of the balances of holdings into sub-balances.
	BalanceBreakdown []*SubBalanceBreakdown1 `xml:"BalBrkdwn,omitempty"`
}

func (f *FinancialInstrumentAggregateBalance2) AddSettledBalance() *FinancialInstrumentQuantity1Choice {
	f.SettledBalance = new(FinancialInstrumentQuantity1Choice)
	return f.SettledBalance
}

func (f *FinancialInstrumentAggregateBalance2) AddTradedBalance() *FinancialInstrumentQuantity1Choice {
	f.TradedBalance = new(FinancialInstrumentQuantity1Choice)
	return f.TradedBalance
}

func (f *FinancialInstrumentAggregateBalance2) AddBalanceBreakdown() *SubBalanceBreakdown1 {
	newValue := new(SubBalanceBreakdown1)
	f.BalanceBreakdown = append(f.BalanceBreakdown, newValue)
	return newValue
}
