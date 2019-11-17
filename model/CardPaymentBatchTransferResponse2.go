package model

// Status of the transactions sent in a previous batch of card payment transactions.
type CardPaymentBatchTransferResponse2 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls,omitempty"`

	// Information related to the previously sent set of transaction.
	DataSet []*CardPaymentDataSet9 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransferResponse2) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransferResponse2) AddDataSet() *CardPaymentDataSet9 {
	newValue := new(CardPaymentDataSet9)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}
