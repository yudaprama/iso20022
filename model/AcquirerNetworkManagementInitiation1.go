package model

// Information related to the network management.
type AcquirerNetworkManagementInitiation1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment6 `xml:"Envt"`

	// Network management transaction.
	Transaction *CardTransaction11 `xml:"Tx"`
}

func (a *AcquirerNetworkManagementInitiation1) AddEnvironment() *CardTransactionEnvironment6 {
	a.Environment = new(CardTransactionEnvironment6)
	return a.Environment
}

func (a *AcquirerNetworkManagementInitiation1) AddTransaction() *CardTransaction11 {
	a.Transaction = new(CardTransaction11)
	return a.Transaction
}
