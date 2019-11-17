package model

// Context in which the card transaction is performed.
type CardTransactionContext3 struct {

	// Context of the card transaction.
	TransactionContext *CardTransactionContext4 `xml:"TxCntxt"`
}

func (c *CardTransactionContext3) AddTransactionContext() *CardTransactionContext4 {
	c.TransactionContext = new(CardTransactionContext4)
	return c.TransactionContext
}
