package model

// Authorisation response from the acquirer.
type AcceptorAuthorisationResponse5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment46 `xml:"Envt"`

	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction53 `xml:"Tx"`

	// Authorisation response from the acquirer.
	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	TransactionResponse *CardPaymentTransaction54 `xml:"TxRspn"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorAuthorisationResponse5) AddEnvironment() *CardPaymentEnvironment46 {
	a.Environment = new(CardPaymentEnvironment46)
	return a.Environment
}

func (a *AcceptorAuthorisationResponse5) AddTransaction() *CardPaymentTransaction53 {
	a.Transaction = new(CardPaymentTransaction53)
	return a.Transaction
}

func (a *AcceptorAuthorisationResponse5) AddTransactionResponse() *CardPaymentTransaction54 {
	a.TransactionResponse = new(CardPaymentTransaction54)
	return a.TransactionResponse
}

func (a *AcceptorAuthorisationResponse5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
