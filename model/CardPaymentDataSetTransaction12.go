package model

// Card payment transaction to be authorised in a batch.
type CardPaymentDataSetTransaction12 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction to authorise.
	Environment *CardPaymentEnvironment32 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext12 `xml:"Cntxt,omitempty"`

	// Card payment transaction to authorise.
	Transaction *CardPaymentTransaction47 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction12) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction12) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction12) AddEnvironment() *CardPaymentEnvironment32 {
	c.Environment = new(CardPaymentEnvironment32)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction12) AddContext() *CardPaymentContext12 {
	c.Context = new(CardPaymentContext12)
	return c.Context
}

func (c *CardPaymentDataSetTransaction12) AddTransaction() *CardPaymentTransaction47 {
	c.Transaction = new(CardPaymentTransaction47)
	return c.Transaction
}
