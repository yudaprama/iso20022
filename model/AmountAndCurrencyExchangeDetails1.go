package model

// Set of elements providing information on the original amount and currency information.
type AmountAndCurrencyExchangeDetails1 struct {

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Reports on currency exchange information.
	CurrencyExchange *CurrencyExchange3 `xml:"CcyXchg,omitempty"`
}

func (a *AmountAndCurrencyExchangeDetails1) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *AmountAndCurrencyExchangeDetails1) AddCurrencyExchange() *CurrencyExchange3 {
	a.CurrencyExchange = new(CurrencyExchange3)
	return a.CurrencyExchange
}
