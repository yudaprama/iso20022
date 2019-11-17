package model

// Set of transactions to capture, sharing common characteristics.
type CardPaymentDataSet4 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Identification of partners involved in the data set building.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification32 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the data set.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls"`

	// Data common to all transactions of the data set.
	CommonData *CommonData2 `xml:"CmonData,omitempty"`

	// Set of transaction to Process.
	Transaction []*CardPaymentDataSetTransaction1Choice `xml:"Tx"`
}

func (c *CardPaymentDataSet4) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet4) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSet4) AddDataSetInitiator() *GenericIdentification32 {
	c.DataSetInitiator = new(GenericIdentification32)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet4) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet4) AddCommonData() *CommonData2 {
	c.CommonData = new(CommonData2)
	return c.CommonData
}

func (c *CardPaymentDataSet4) AddTransaction() *CardPaymentDataSetTransaction1Choice {
	newValue := new(CardPaymentDataSetTransaction1Choice)
	c.Transaction = append(c.Transaction, newValue)
	return newValue
}
