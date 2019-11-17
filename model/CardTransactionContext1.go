package model

// Context in which the card transaction is performed.
type CardTransactionContext1 struct {

	// Context of the card transaction.
	TransactionContext *CardTransactionContext2 `xml:"TxCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardTransactionContext1) AddTransactionContext() *CardTransactionContext2 {
	c.TransactionContext = new(CardTransactionContext2)
	return c.TransactionContext
}

func (c *CardTransactionContext1) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
