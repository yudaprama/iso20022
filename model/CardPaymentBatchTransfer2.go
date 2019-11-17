package model

// Card payment transactions from one or several data set of transactions.
type CardPaymentBatchTransfer2 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls,omitempty"`

	// Card payment transactions from one data set of transactions.
	DataSet []*CardPaymentDataSet7 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransfer2) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransfer2) AddDataSet() *CardPaymentDataSet7 {
	newValue := new(CardPaymentDataSet7)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}
