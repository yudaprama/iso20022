package model

// Cancelled card payment transaction to be captured in a batch.
type CardPaymentDataSetTransaction3 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the cancelled transaction captured in batch.
	Environment *CardPaymentEnvironment14 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext4 `xml:"Cntxt,omitempty"`

	// Card payment cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction20 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction3) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction3) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction3) AddEnvironment() *CardPaymentEnvironment14 {
	c.Environment = new(CardPaymentEnvironment14)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction3) AddContext() *CardPaymentContext4 {
	c.Context = new(CardPaymentContext4)
	return c.Context
}

func (c *CardPaymentDataSetTransaction3) AddTransaction() *CardPaymentTransaction20 {
	c.Transaction = new(CardPaymentTransaction20)
	return c.Transaction
}
