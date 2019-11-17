package model

// Currency conversion outcome from the service provider.
type AcceptorCurrencyConversionResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Currency conversion of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction38 `xml:"Tx"`

	// Details of the currency conversion.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs"`
}

func (a *AcceptorCurrencyConversionResponse2) AddEnvironment() *CardPaymentEnvironment33 {
	a.Environment = new(CardPaymentEnvironment33)
	return a.Environment
}

func (a *AcceptorCurrencyConversionResponse2) AddTransaction() *CardPaymentTransaction38 {
	a.Transaction = new(CardPaymentTransaction38)
	return a.Transaction
}

func (a *AcceptorCurrencyConversionResponse2) AddCurrencyConversion() *CurrencyConversion3 {
	a.CurrencyConversion = new(CurrencyConversion3)
	return a.CurrencyConversion
}
