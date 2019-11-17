package model

// Set of elements providing information on the original amount and currency information.
type AmountAndCurrencyExchangeDetails2 struct {

	// Identifies the type of amount.
	Type *Max35Text `xml:"Tp"`

	// Identifies the proprietary amount.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Reports on currency exchange information.
	CurrencyExchange *CurrencyExchange3 `xml:"CcyXchg,omitempty"`
}

func (a *AmountAndCurrencyExchangeDetails2) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *AmountAndCurrencyExchangeDetails2) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *AmountAndCurrencyExchangeDetails2) AddCurrencyExchange() *CurrencyExchange3 {
	a.CurrencyExchange = new(CurrencyExchange3)
	return a.CurrencyExchange
}
