package model

// Cancelled card payment transaction to be captured in a batch.
type CardPaymentDataSetTransaction15 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the cancelled transaction captured in batch.
	Environment *CardPaymentEnvironment52 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext19 `xml:"Cntxt,omitempty"`

	// Card payment cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction61 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction15) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction15) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction15) AddEnvironment() *CardPaymentEnvironment52 {
	c.Environment = new(CardPaymentEnvironment52)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction15) AddContext() *CardPaymentContext19 {
	c.Context = new(CardPaymentContext19)
	return c.Context
}

func (c *CardPaymentDataSetTransaction15) AddTransaction() *CardPaymentTransaction61 {
	c.Transaction = new(CardPaymentTransaction61)
	return c.Transaction
}
