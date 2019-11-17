package model

// Card payment transactions from one or several data set of transactions.
type CardPaymentBatchTransfer1 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls,omitempty"`

	// Card payment transactions from one data set of transactions.
	DataSet []*CardPaymentDataSet4 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransfer1) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransfer1) AddDataSet() *CardPaymentDataSet4 {
	newValue := new(CardPaymentDataSet4)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}
