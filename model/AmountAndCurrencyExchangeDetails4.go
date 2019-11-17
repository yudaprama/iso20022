package model

// Set of elements used to provide information on the original amount and currency exchange.
type AmountAndCurrencyExchangeDetails4 struct {

	// Specifies the type of amount.
	Type *Max35Text `xml:"Tp"`

	// Amount of money to be exchanged against another amount of money in the counter currency.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Set of elements used to provide details on the currency exchange.
	CurrencyExchange *CurrencyExchange5 `xml:"CcyXchg,omitempty"`
}

func (a *AmountAndCurrencyExchangeDetails4) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *AmountAndCurrencyExchangeDetails4) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndCurrencyExchangeDetails4) AddCurrencyExchange() *CurrencyExchange5 {
	a.CurrencyExchange = new(CurrencyExchange5)
	return a.CurrencyExchange
}
