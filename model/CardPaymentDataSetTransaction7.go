package model

// Cancelled card payment transaction to be captured in a batch.
type CardPaymentDataSetTransaction7 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the cancelled transaction captured in batch.
	Environment *CardPaymentEnvironment27 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext4 `xml:"Cntxt,omitempty"`

	// Card payment cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction30 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction7) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction7) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction7) AddEnvironment() *CardPaymentEnvironment27 {
	c.Environment = new(CardPaymentEnvironment27)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction7) AddContext() *CardPaymentContext4 {
	c.Context = new(CardPaymentContext4)
	return c.Context
}

func (c *CardPaymentDataSetTransaction7) AddTransaction() *CardPaymentTransaction30 {
	c.Transaction = new(CardPaymentTransaction30)
	return c.Transaction
}
