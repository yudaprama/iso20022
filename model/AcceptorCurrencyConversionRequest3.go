package model

// Information related to the currency conversion request.
type AcceptorCurrencyConversionRequest3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment56 `xml:"Envt"`

	// Card payment transaction for which the currency conversion is requested.
	Transaction *CardPaymentTransaction64 `xml:"Tx"`
}

func (a *AcceptorCurrencyConversionRequest3) AddEnvironment() *CardPaymentEnvironment56 {
	a.Environment = new(CardPaymentEnvironment56)
	return a.Environment
}

func (a *AcceptorCurrencyConversionRequest3) AddTransaction() *CardPaymentTransaction64 {
	a.Transaction = new(CardPaymentTransaction64)
	return a.Transaction
}
