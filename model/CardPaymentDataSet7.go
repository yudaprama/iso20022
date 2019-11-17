package model

// Set of transactions to capture, sharing common characteristics.
type CardPaymentDataSet7 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Identification of partners involved in the data set building.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification32 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the data set.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls"`

	// Data common to all transactions of the data set.
	CommonData *CommonData3 `xml:"CmonData,omitempty"`

	// Set of transaction to Process.
	Transaction []*CardPaymentDataSetTransaction2Choice `xml:"Tx"`
}

func (c *CardPaymentDataSet7) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet7) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSet7) AddDataSetInitiator() *GenericIdentification32 {
	c.DataSetInitiator = new(GenericIdentification32)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet7) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet7) AddCommonData() *CommonData3 {
	c.CommonData = new(CommonData3)
	return c.CommonData
}

func (c *CardPaymentDataSet7) AddTransaction() *CardPaymentDataSetTransaction2Choice {
	newValue := new(CardPaymentDataSetTransaction2Choice)
	c.Transaction = append(c.Transaction, newValue)
	return newValue
}
