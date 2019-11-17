package model

// Card payment transaction to be authorised in a batch.
type CardPaymentDataSetTransaction4 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction to authorise.
	Environment *CardPaymentEnvironment14 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext4 `xml:"Cntxt,omitempty"`

	// Card payment transaction to authorise.
	Transaction *CardPaymentTransaction19 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction4) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction4) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction4) AddEnvironment() *CardPaymentEnvironment14 {
	c.Environment = new(CardPaymentEnvironment14)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction4) AddContext() *CardPaymentContext4 {
	c.Context = new(CardPaymentContext4)
	return c.Context
}

func (c *CardPaymentDataSetTransaction4) AddTransaction() *CardPaymentTransaction19 {
	c.Transaction = new(CardPaymentTransaction19)
	return c.Transaction
}
