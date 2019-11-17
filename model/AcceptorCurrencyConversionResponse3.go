package model

// Currency conversion outcome from the service provider.
type AcceptorCurrencyConversionResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment46 `xml:"Envt"`

	// Currency conversion of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction53 `xml:"Tx"`

	// Details of the currency conversion.
	CurrencyConversionResult *CurrencyConversion7 `xml:"CcyConvsRslt"`
}

func (a *AcceptorCurrencyConversionResponse3) AddEnvironment() *CardPaymentEnvironment46 {
	a.Environment = new(CardPaymentEnvironment46)
	return a.Environment
}

func (a *AcceptorCurrencyConversionResponse3) AddTransaction() *CardPaymentTransaction53 {
	a.Transaction = new(CardPaymentTransaction53)
	return a.Transaction
}

func (a *AcceptorCurrencyConversionResponse3) AddCurrencyConversionResult() *CurrencyConversion7 {
	a.CurrencyConversionResult = new(CurrencyConversion7)
	return a.CurrencyConversionResult
}
