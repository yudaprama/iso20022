package model

// Information related to the response to a key exchange.
type AcquirerKeyExchangeResponse1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment6 `xml:"Envt"`

	// Key exchange transaction.
	Transaction *CardTransaction14 `xml:"Tx"`
}

func (a *AcquirerKeyExchangeResponse1) AddEnvironment() *CardTransactionEnvironment6 {
	a.Environment = new(CardTransactionEnvironment6)
	return a.Environment
}

func (a *AcquirerKeyExchangeResponse1) AddTransaction() *CardTransaction14 {
	a.Transaction = new(CardTransaction14)
	return a.Transaction
}
