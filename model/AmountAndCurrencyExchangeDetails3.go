package model

// Set of elements used to provide information on the original amount and currency exchange.
type AmountAndCurrencyExchangeDetails3 struct {

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Set of elements used to provide details on the currency exchange.
	CurrencyExchange *CurrencyExchange5 `xml:"CcyXchg,omitempty"`
}

func (a *AmountAndCurrencyExchangeDetails3) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndCurrencyExchangeDetails3) AddCurrencyExchange() *CurrencyExchange5 {
	a.CurrencyExchange = new(CurrencyExchange5)
	return a.CurrencyExchange
}
