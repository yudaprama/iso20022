package model

// Authorisation response from the acquirer.
type AcceptorAuthorisationResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction38 `xml:"Tx"`

	// Authorisation response from the acquirer.
	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	TransactionResponse *CardPaymentTransaction39 `xml:"TxRspn"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorAuthorisationResponse4) AddEnvironment() *CardPaymentEnvironment33 {
	a.Environment = new(CardPaymentEnvironment33)
	return a.Environment
}

func (a *AcceptorAuthorisationResponse4) AddTransaction() *CardPaymentTransaction38 {
	a.Transaction = new(CardPaymentTransaction38)
	return a.Transaction
}

func (a *AcceptorAuthorisationResponse4) AddTransactionResponse() *CardPaymentTransaction39 {
	a.TransactionResponse = new(CardPaymentTransaction39)
	return a.TransactionResponse
}

func (a *AcceptorAuthorisationResponse4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
