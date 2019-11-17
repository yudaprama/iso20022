package model

// Authorisation response from the acquirer.
type AcceptorAuthorisationResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment3 `xml:"Envt"`

	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction2 `xml:"Tx"`

	// Authorisation response from the acquirer.
	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	TransactionResponse *CardPaymentTransaction9 `xml:"TxRspn"`
}

func (a *AcceptorAuthorisationResponse1) AddEnvironment() *CardPaymentEnvironment3 {
	a.Environment = new(CardPaymentEnvironment3)
	return a.Environment
}

func (a *AcceptorAuthorisationResponse1) AddTransaction() *CardPaymentTransaction2 {
	a.Transaction = new(CardPaymentTransaction2)
	return a.Transaction
}

func (a *AcceptorAuthorisationResponse1) AddTransactionResponse() *CardPaymentTransaction9 {
	a.TransactionResponse = new(CardPaymentTransaction9)
	return a.TransactionResponse
}
