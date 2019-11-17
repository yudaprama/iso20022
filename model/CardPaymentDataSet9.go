package model

// Result of the captured set of transactions.
type CardPaymentDataSet9 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Result of the data set capture.
	DataSetResult *ResponseType1 `xml:"DataSetRslt"`

	// Indicates if the data set must be removed from the POI (Point Of Interaction).
	RemoveDataSet *TrueFalseIndicator `xml:"RmvDataSet"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification32 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the batch.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls"`

	// Transaction in the batch, whose capture has been rejected.
	RejectedTransaction []*CardPaymentDataSet8 `xml:"RjctdTx,omitempty"`
}

func (c *CardPaymentDataSet9) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet9) AddDataSetResult() *ResponseType1 {
	c.DataSetResult = new(ResponseType1)
	return c.DataSetResult
}

func (c *CardPaymentDataSet9) SetRemoveDataSet(value string) {
	c.RemoveDataSet = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentDataSet9) AddDataSetInitiator() *GenericIdentification32 {
	c.DataSetInitiator = new(GenericIdentification32)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet9) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet9) AddRejectedTransaction() *CardPaymentDataSet8 {
	newValue := new(CardPaymentDataSet8)
	c.RejectedTransaction = append(c.RejectedTransaction, newValue)
	return newValue
}
