package model

// Currencey conversion outcome from the service provider.
type AcceptorCurrencyConversionResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Currency conversion of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction23 `xml:"Tx"`

	// Result of the currency conversion.
	TransactionResponse *Response1Code `xml:"TxRspn"`

	// Details of the currency conversion.
	CurrencyConversion *CurrencyConversion1 `xml:"CcyConvs"`
}

func (a *AcceptorCurrencyConversionResponse1) AddEnvironment() *CardPaymentEnvironment21 {
	a.Environment = new(CardPaymentEnvironment21)
	return a.Environment
}

func (a *AcceptorCurrencyConversionResponse1) AddTransaction() *CardPaymentTransaction23 {
	a.Transaction = new(CardPaymentTransaction23)
	return a.Transaction
}

func (a *AcceptorCurrencyConversionResponse1) SetTransactionResponse(value string) {
	a.TransactionResponse = (*Response1Code)(&value)
}

func (a *AcceptorCurrencyConversionResponse1) AddCurrencyConversion() *CurrencyConversion1 {
	a.CurrencyConversion = new(CurrencyConversion1)
	return a.CurrencyConversion
}
