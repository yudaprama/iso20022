package model

// Transaction to capture in the batch.
type CardPaymentDataSetTransaction1 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the transaction in a transaction captured in batch.
	Environment *CardPaymentEnvironment6 `xml:"Envt"`

	// Data related to the context of the transaction.
	Context *CardPaymentContext3 `xml:"Cntxt,omitempty"`

	// Transaction information to be captured.
	Transaction *CardPaymentTransaction4 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction1) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction1) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction1) AddEnvironment() *CardPaymentEnvironment6 {
	c.Environment = new(CardPaymentEnvironment6)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction1) AddContext() *CardPaymentContext3 {
	c.Context = new(CardPaymentContext3)
	return c.Context
}

func (c *CardPaymentDataSetTransaction1) AddTransaction() *CardPaymentTransaction4 {
	c.Transaction = new(CardPaymentTransaction4)
	return c.Transaction
}
