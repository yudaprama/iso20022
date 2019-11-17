package model

// Card payment transactions from one or several data set of transactions.
type CardPaymentBatchTransfer4 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals7 `xml:"TxTtls,omitempty"`

	// Card payment transactions from one data set of transactions.
	DataSet []*CardPaymentDataSet13 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransfer4) AddTransactionTotals() *TransactionTotals7 {
	newValue := new(TransactionTotals7)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransfer4) AddDataSet() *CardPaymentDataSet13 {
	newValue := new(CardPaymentDataSet13)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}
