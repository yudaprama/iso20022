package model

// Status of the transactions sent in a previous batch of card payment transactions.
type CardPaymentBatchTransferResponse3 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals3 `xml:"TxTtls,omitempty"`

	// Information related to the previously sent set of transaction.
	DataSet []*CardPaymentDataSet12 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransferResponse3) AddTransactionTotals() *TransactionTotals3 {
	newValue := new(TransactionTotals3)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransferResponse3) AddDataSet() *CardPaymentDataSet12 {
	newValue := new(CardPaymentDataSet12)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}
