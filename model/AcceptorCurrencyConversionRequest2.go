package model

// Information related to the currency conversion request.
type AcceptorCurrencyConversionRequest2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment44 `xml:"Envt"`

	// Card payment transaction for which the currency conversion is requested.
	Transaction *CardPaymentTransaction49 `xml:"Tx"`
}

func (a *AcceptorCurrencyConversionRequest2) AddEnvironment() *CardPaymentEnvironment44 {
	a.Environment = new(CardPaymentEnvironment44)
	return a.Environment
}

func (a *AcceptorCurrencyConversionRequest2) AddTransaction() *CardPaymentTransaction49 {
	a.Transaction = new(CardPaymentTransaction49)
	return a.Transaction
}
