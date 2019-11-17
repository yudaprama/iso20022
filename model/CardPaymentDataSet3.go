package model

// Transaction for whose batch capture has been rejected.
type CardPaymentDataSet3 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Response to the capture of the transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Data related to the environment of the transaction.
	Environment *CardPaymentEnvironment3 `xml:"Envt"`

	// Transaction that has been rejected.
	Transaction *CardPaymentTransactionAdviceResponse1 `xml:"Tx"`
}

func (c *CardPaymentDataSet3) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSet3) AddTransactionResponse() *ResponseType1 {
	c.TransactionResponse = new(ResponseType1)
	return c.TransactionResponse
}

func (c *CardPaymentDataSet3) AddEnvironment() *CardPaymentEnvironment3 {
	c.Environment = new(CardPaymentEnvironment3)
	return c.Environment
}

func (c *CardPaymentDataSet3) AddTransaction() *CardPaymentTransactionAdviceResponse1 {
	c.Transaction = new(CardPaymentTransactionAdviceResponse1)
	return c.Transaction
}
