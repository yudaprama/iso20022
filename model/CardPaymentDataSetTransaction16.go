package model

// Card payment transaction to be authorised in a batch.
type CardPaymentDataSetTransaction16 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction to authorise.
	Environment *CardPaymentEnvironment53 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext18 `xml:"Cntxt,omitempty"`

	// Card payment transaction to authorise.
	Transaction *CardPaymentTransaction62 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction16) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction16) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction16) AddEnvironment() *CardPaymentEnvironment53 {
	c.Environment = new(CardPaymentEnvironment53)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction16) AddContext() *CardPaymentContext18 {
	c.Context = new(CardPaymentContext18)
	return c.Context
}

func (c *CardPaymentDataSetTransaction16) AddTransaction() *CardPaymentTransaction62 {
	c.Transaction = new(CardPaymentTransaction62)
	return c.Transaction
}
