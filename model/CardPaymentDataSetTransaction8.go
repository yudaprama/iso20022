package model

// Card payment transaction to be authorised in a batch.
type CardPaymentDataSetTransaction8 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction to authorise.
	Environment *CardPaymentEnvironment27 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext7 `xml:"Cntxt,omitempty"`

	// Card payment transaction to authorise.
	Transaction *CardPaymentTransaction31 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction8) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction8) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction8) AddEnvironment() *CardPaymentEnvironment27 {
	c.Environment = new(CardPaymentEnvironment27)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction8) AddContext() *CardPaymentContext7 {
	c.Context = new(CardPaymentContext7)
	return c.Context
}

func (c *CardPaymentDataSetTransaction8) AddTransaction() *CardPaymentTransaction31 {
	c.Transaction = new(CardPaymentTransaction31)
	return c.Transaction
}
