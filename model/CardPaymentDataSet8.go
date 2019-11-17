package model

// Transaction for whose batch capture has been rejected.
type CardPaymentDataSet8 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Response to the capture of the transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Data related to the environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Transaction that has been rejected.
	Transaction *CardPaymentTransactionAdviceResponse5 `xml:"Tx"`
}

func (c *CardPaymentDataSet8) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSet8) AddTransactionResponse() *ResponseType1 {
	c.TransactionResponse = new(ResponseType1)
	return c.TransactionResponse
}

func (c *CardPaymentDataSet8) AddEnvironment() *CardPaymentEnvironment21 {
	c.Environment = new(CardPaymentEnvironment21)
	return c.Environment
}

func (c *CardPaymentDataSet8) AddTransaction() *CardPaymentTransactionAdviceResponse5 {
	c.Transaction = new(CardPaymentTransactionAdviceResponse5)
	return c.Transaction
}
