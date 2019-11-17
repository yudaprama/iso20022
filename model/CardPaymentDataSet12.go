package model

// Result of the captured set of transactions.
type CardPaymentDataSet12 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Result of the data set capture.
	DataSetResult *ResponseType1 `xml:"DataSetRslt"`

	// Indicates if the data set must be removed from the POI (Point Of Interaction).
	RemoveDataSet *TrueFalseIndicator `xml:"RmvDataSet"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification53 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the batch.
	TransactionTotals []*TransactionTotals3 `xml:"TxTtls"`

	// Transaction in the batch, whose capture has been rejected.
	RejectedTransaction []*CardPaymentDataSet11 `xml:"RjctdTx,omitempty"`
}

func (c *CardPaymentDataSet12) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet12) AddDataSetResult() *ResponseType1 {
	c.DataSetResult = new(ResponseType1)
	return c.DataSetResult
}

func (c *CardPaymentDataSet12) SetRemoveDataSet(value string) {
	c.RemoveDataSet = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentDataSet12) AddDataSetInitiator() *GenericIdentification53 {
	c.DataSetInitiator = new(GenericIdentification53)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet12) AddTransactionTotals() *TransactionTotals3 {
	newValue := new(TransactionTotals3)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet12) AddRejectedTransaction() *CardPaymentDataSet11 {
	newValue := new(CardPaymentDataSet11)
	c.RejectedTransaction = append(c.RejectedTransaction, newValue)
	return newValue
}
