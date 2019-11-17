package model

// Cancelled card payment transaction to be captured in a batch.
type CardPaymentDataSetTransaction11 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the cancelled transaction captured in batch.
	Environment *CardPaymentEnvironment40 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext13 `xml:"Cntxt,omitempty"`

	// Card payment cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction46 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction11) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction11) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction11) AddEnvironment() *CardPaymentEnvironment40 {
	c.Environment = new(CardPaymentEnvironment40)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction11) AddContext() *CardPaymentContext13 {
	c.Context = new(CardPaymentContext13)
	return c.Context
}

func (c *CardPaymentDataSetTransaction11) AddTransaction() *CardPaymentTransaction46 {
	c.Transaction = new(CardPaymentTransaction46)
	return c.Transaction
}
