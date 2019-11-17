package model

// Status of the transactions sent in a previous batch of card payment transactions.
type CardPaymentBatchTransferResponse4 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals7 `xml:"TxTtls,omitempty"`

	// Information related to the previously sent set of transaction.
	DataSet []*CardPaymentDataSet14 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransferResponse4) AddTransactionTotals() *TransactionTotals7 {
	newValue := new(TransactionTotals7)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransferResponse4) AddDataSet() *CardPaymentDataSet14 {
	newValue := new(CardPaymentDataSet14)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}
