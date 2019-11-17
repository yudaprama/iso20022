package model

// Completed card payment transaction to be captured in batch.
type CardPaymentDataSetTransaction14 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the transaction in a transaction captured in batch.
	Environment *CardPaymentEnvironment52 `xml:"Envt"`

	// Data related to the context of the transaction.
	Context *CardPaymentContext18 `xml:"Cntxt,omitempty"`

	// Transaction information to be captured.
	Transaction *CardPaymentTransaction60 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction14) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction14) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction14) AddEnvironment() *CardPaymentEnvironment52 {
	c.Environment = new(CardPaymentEnvironment52)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction14) AddContext() *CardPaymentContext18 {
	c.Context = new(CardPaymentContext18)
	return c.Context
}

func (c *CardPaymentDataSetTransaction14) AddTransaction() *CardPaymentTransaction60 {
	c.Transaction = new(CardPaymentTransaction60)
	return c.Transaction
}
