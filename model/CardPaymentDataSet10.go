package model

// Set of transactions to capture, sharing common characteristics.
type CardPaymentDataSet10 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Identification of partners involved in the data set building.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification53 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the data set.
	TransactionTotals []*TransactionTotals3 `xml:"TxTtls"`

	// Data common to all transactions of the data set.
	CommonData *CommonData4 `xml:"CmonData,omitempty"`

	// Set of transaction to Process.
	Transaction []*CardPaymentDataSetTransaction3Choice `xml:"Tx"`
}

func (c *CardPaymentDataSet10) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet10) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSet10) AddDataSetInitiator() *GenericIdentification53 {
	c.DataSetInitiator = new(GenericIdentification53)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet10) AddTransactionTotals() *TransactionTotals3 {
	newValue := new(TransactionTotals3)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet10) AddCommonData() *CommonData4 {
	c.CommonData = new(CommonData4)
	return c.CommonData
}

func (c *CardPaymentDataSet10) AddTransaction() *CardPaymentDataSetTransaction3Choice {
	newValue := new(CardPaymentDataSetTransaction3Choice)
	c.Transaction = append(c.Transaction, newValue)
	return newValue
}
