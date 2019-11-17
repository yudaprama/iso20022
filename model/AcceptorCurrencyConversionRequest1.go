package model

// Information related to the currency conversion request.
type AcceptorCurrencyConversionRequest1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment30 `xml:"Envt"`

	// Card payment transaction for which the currency conversion is requested.
	Transaction *CardPaymentTransaction34 `xml:"Tx"`
}

func (a *AcceptorCurrencyConversionRequest1) AddEnvironment() *CardPaymentEnvironment30 {
	a.Environment = new(CardPaymentEnvironment30)
	return a.Environment
}

func (a *AcceptorCurrencyConversionRequest1) AddTransaction() *CardPaymentTransaction34 {
	a.Transaction = new(CardPaymentTransaction34)
	return a.Transaction
}
