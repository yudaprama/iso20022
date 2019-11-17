package model

// Set of transactions to capture, sharing common characteristics.
type CardPaymentDataSet13 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification5 `xml:"DataSetId"`

	// Identification of partners involved in the data set building.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification53 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the data set.
	TransactionTotals []*TransactionTotals7 `xml:"TxTtls"`

	// Data common to all transactions of the data set.
	CommonData *CommonData5 `xml:"CmonData,omitempty"`

	// Set of transaction to Process.
	Transaction []*CardPaymentDataSetTransaction4Choice `xml:"Tx"`
}

func (c *CardPaymentDataSet13) AddDataSetIdentification() *DataSetIdentification5 {
	c.DataSetIdentification = new(DataSetIdentification5)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet13) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSet13) AddDataSetInitiator() *GenericIdentification53 {
	c.DataSetInitiator = new(GenericIdentification53)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet13) AddTransactionTotals() *TransactionTotals7 {
	newValue := new(TransactionTotals7)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet13) AddCommonData() *CommonData5 {
	c.CommonData = new(CommonData5)
	return c.CommonData
}

func (c *CardPaymentDataSet13) AddTransaction() *CardPaymentDataSetTransaction4Choice {
	newValue := new(CardPaymentDataSetTransaction4Choice)
	c.Transaction = append(c.Transaction, newValue)
	return newValue
}
