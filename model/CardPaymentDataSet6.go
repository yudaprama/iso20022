package model

// Transaction for whose batch capture has been rejected.
type CardPaymentDataSet6 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Response to the capture of the transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Data related to the environment of the transaction.
	Environment *CardPaymentEnvironment11 `xml:"Envt"`

	// Transaction that has been rejected.
	Transaction *CardPaymentTransactionAdviceResponse1 `xml:"Tx"`
}

func (c *CardPaymentDataSet6) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSet6) AddTransactionResponse() *ResponseType1 {
	c.TransactionResponse = new(ResponseType1)
	return c.TransactionResponse
}

func (c *CardPaymentDataSet6) AddEnvironment() *CardPaymentEnvironment11 {
	c.Environment = new(CardPaymentEnvironment11)
	return c.Environment
}

func (c *CardPaymentDataSet6) AddTransaction() *CardPaymentTransactionAdviceResponse1 {
	c.Transaction = new(CardPaymentTransactionAdviceResponse1)
	return c.Transaction
}
