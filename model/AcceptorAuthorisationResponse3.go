package model

// Authorisation response from the acquirer.
type AcceptorAuthorisationResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction23 `xml:"Tx"`

	// Authorisation response from the acquirer.
	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	TransactionResponse *CardPaymentTransaction24 `xml:"TxRspn"`
}

func (a *AcceptorAuthorisationResponse3) AddEnvironment() *CardPaymentEnvironment21 {
	a.Environment = new(CardPaymentEnvironment21)
	return a.Environment
}

func (a *AcceptorAuthorisationResponse3) AddTransaction() *CardPaymentTransaction23 {
	a.Transaction = new(CardPaymentTransaction23)
	return a.Transaction
}

func (a *AcceptorAuthorisationResponse3) AddTransactionResponse() *CardPaymentTransaction24 {
	a.TransactionResponse = new(CardPaymentTransaction24)
	return a.TransactionResponse
}
