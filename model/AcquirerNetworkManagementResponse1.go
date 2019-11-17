package model

// Information related to the response to the network management.
type AcquirerNetworkManagementResponse1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment6 `xml:"Envt"`

	// Network management transaction.
	Transaction *CardTransaction12 `xml:"Tx"`
}

func (a *AcquirerNetworkManagementResponse1) AddEnvironment() *CardTransactionEnvironment6 {
	a.Environment = new(CardTransactionEnvironment6)
	return a.Environment
}

func (a *AcquirerNetworkManagementResponse1) AddTransaction() *CardTransaction12 {
	a.Transaction = new(CardTransaction12)
	return a.Transaction
}
