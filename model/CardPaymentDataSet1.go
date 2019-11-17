package model

// Set of transactions to capture, sharing common characteristics.
type CardPaymentDataSet1 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Identification of partners involved in the data set building.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification32 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the data set.
	TransactionTotals []*TransactionTotals1 `xml:"TxTtls"`

	// Data common to all transactions of the data set.
	CommonData *CommonData1 `xml:"CmonData,omitempty"`

	// Set of transaction to capture.
	TransactionToCapture []*CardPaymentDataSetTransaction1 `xml:"TxToCaptr,omitempty"`
}

func (c *CardPaymentDataSet1) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet1) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSet1) AddDataSetInitiator() *GenericIdentification32 {
	c.DataSetInitiator = new(GenericIdentification32)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet1) AddTransactionTotals() *TransactionTotals1 {
	newValue := new(TransactionTotals1)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet1) AddCommonData() *CommonData1 {
	c.CommonData = new(CommonData1)
	return c.CommonData
}

func (c *CardPaymentDataSet1) AddTransactionToCapture() *CardPaymentDataSetTransaction1 {
	newValue := new(CardPaymentDataSetTransaction1)
	c.TransactionToCapture = append(c.TransactionToCapture, newValue)
	return newValue
}
