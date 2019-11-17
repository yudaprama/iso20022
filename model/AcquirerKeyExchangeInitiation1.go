package model

// Information related to the key exchange.
type AcquirerKeyExchangeInitiation1 struct {

	// Environment of the transaction.
	Environment *CardTransactionEnvironment6 `xml:"Envt"`

	// Key exchange transaction.
	Transaction *CardTransaction13 `xml:"Tx"`
}

func (a *AcquirerKeyExchangeInitiation1) AddEnvironment() *CardTransactionEnvironment6 {
	a.Environment = new(CardTransactionEnvironment6)
	return a.Environment
}

func (a *AcquirerKeyExchangeInitiation1) AddTransaction() *CardTransaction13 {
	a.Transaction = new(CardTransaction13)
	return a.Transaction
}
