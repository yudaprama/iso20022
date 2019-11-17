package model

// Authorisation response from the acquirer.
type AcceptorAuthorisationResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment11 `xml:"Envt"`

	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction2 `xml:"Tx"`

	// Authorisation response from the acquirer.
	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	TransactionResponse *CardPaymentTransaction18 `xml:"TxRspn"`
}

func (a *AcceptorAuthorisationResponse2) AddEnvironment() *CardPaymentEnvironment11 {
	a.Environment = new(CardPaymentEnvironment11)
	return a.Environment
}

func (a *AcceptorAuthorisationResponse2) AddTransaction() *CardPaymentTransaction2 {
	a.Transaction = new(CardPaymentTransaction2)
	return a.Transaction
}

func (a *AcceptorAuthorisationResponse2) AddTransactionResponse() *CardPaymentTransaction18 {
	a.TransactionResponse = new(CardPaymentTransaction18)
	return a.TransactionResponse
}
