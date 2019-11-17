package model

// Status of the transactions sent in a previous batch of card payment transactions.
type CardPaymentBatchTransferResponse1 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals *TransactionTotals2 `xml:"TxTtls"`

	// Information related to the previously sent set of transaction.
	DataSet []*CardPaymentDataSet5 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransferResponse1) AddTransactionTotals() *TransactionTotals2 {
	c.TransactionTotals = new(TransactionTotals2)
	return c.TransactionTotals
}

func (c *CardPaymentBatchTransferResponse1) AddDataSet() *CardPaymentDataSet5 {
	newValue := new(CardPaymentDataSet5)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}
