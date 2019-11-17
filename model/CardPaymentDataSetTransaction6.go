package model

// Completed card payment transaction to be captured in batch.
type CardPaymentDataSetTransaction6 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the transaction in a transaction captured in batch.
	Environment *CardPaymentEnvironment27 `xml:"Envt"`

	// Data related to the context of the transaction.
	Context *CardPaymentContext7 `xml:"Cntxt,omitempty"`

	// Transaction information to be captured.
	Transaction *CardPaymentTransaction29 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction6) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction6) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction6) AddEnvironment() *CardPaymentEnvironment27 {
	c.Environment = new(CardPaymentEnvironment27)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction6) AddContext() *CardPaymentContext7 {
	c.Context = new(CardPaymentContext7)
	return c.Context
}

func (c *CardPaymentDataSetTransaction6) AddTransaction() *CardPaymentTransaction29 {
	c.Transaction = new(CardPaymentTransaction29)
	return c.Transaction
}
