package model

// Result of the captured set of transactions.
type CardPaymentDataSet14 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification5 `xml:"DataSetId"`

	// Result of the data set capture.
	DataSetResult *ResponseType5 `xml:"DataSetRslt"`

	// Indicates if the data set must be removed from the POI (Point Of Interaction).
	RemoveDataSet *TrueFalseIndicator `xml:"RmvDataSet"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification53 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the batch.
	TransactionTotals []*TransactionTotals7 `xml:"TxTtls"`

	// Transaction in the batch, whose capture has been rejected.
	RejectedTransaction []*CardPaymentDataSet15 `xml:"RjctdTx,omitempty"`
}

func (c *CardPaymentDataSet14) AddDataSetIdentification() *DataSetIdentification5 {
	c.DataSetIdentification = new(DataSetIdentification5)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet14) AddDataSetResult() *ResponseType5 {
	c.DataSetResult = new(ResponseType5)
	return c.DataSetResult
}

func (c *CardPaymentDataSet14) SetRemoveDataSet(value string) {
	c.RemoveDataSet = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentDataSet14) AddDataSetInitiator() *GenericIdentification53 {
	c.DataSetInitiator = new(GenericIdentification53)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet14) AddTransactionTotals() *TransactionTotals7 {
	newValue := new(TransactionTotals7)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet14) AddRejectedTransaction() *CardPaymentDataSet15 {
	newValue := new(CardPaymentDataSet15)
	c.RejectedTransaction = append(c.RejectedTransaction, newValue)
	return newValue
}
