package model

// Completed card payment transaction to be captured in batch.
type CardPaymentDataSetTransaction2 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the transaction in a transaction captured in batch.
	Environment *CardPaymentEnvironment14 `xml:"Envt"`

	// Data related to the context of the transaction.
	Context *CardPaymentContext4 `xml:"Cntxt,omitempty"`

	// Transaction information to be captured.
	Transaction *CardPaymentTransaction14 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction2) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction2) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction2) AddEnvironment() *CardPaymentEnvironment14 {
	c.Environment = new(CardPaymentEnvironment14)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction2) AddContext() *CardPaymentContext4 {
	c.Context = new(CardPaymentContext4)
	return c.Context
}

func (c *CardPaymentDataSetTransaction2) AddTransaction() *CardPaymentTransaction14 {
	c.Transaction = new(CardPaymentTransaction14)
	return c.Transaction
}
