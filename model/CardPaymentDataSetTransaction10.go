package model

// Completed card payment transaction to be captured in batch.
type CardPaymentDataSetTransaction10 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the transaction in a transaction captured in batch.
	Environment *CardPaymentEnvironment40 `xml:"Envt"`

	// Data related to the context of the transaction.
	Context *CardPaymentContext12 `xml:"Cntxt,omitempty"`

	// Transaction information to be captured.
	Transaction *CardPaymentTransaction45 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction10) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction10) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction10) AddEnvironment() *CardPaymentEnvironment40 {
	c.Environment = new(CardPaymentEnvironment40)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction10) AddContext() *CardPaymentContext12 {
	c.Context = new(CardPaymentContext12)
	return c.Context
}

func (c *CardPaymentDataSetTransaction10) AddTransaction() *CardPaymentTransaction45 {
	c.Transaction = new(CardPaymentTransaction45)
	return c.Transaction
}
