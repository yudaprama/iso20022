package model

// Card payment transactions from one or several data set of transactions.
type CardPaymentBatchTransfer3 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals3 `xml:"TxTtls,omitempty"`

	// Card payment transactions from one data set of transactions.
	DataSet []*CardPaymentDataSet10 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransfer3) AddTransactionTotals() *TransactionTotals3 {
	newValue := new(TransactionTotals3)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransfer3) AddDataSet() *CardPaymentDataSet10 {
	newValue := new(CardPaymentDataSet10)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}
