package model

// Card payment transaction including an authorisation response.
type CardPaymentDataSetTransaction5 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction.
	Environment *CardPaymentEnvironment16 `xml:"Envt"`

	// Card payment transaction authorisation result.
	Transaction *CardPaymentTransaction12 `xml:"Tx"`

	// Response to the authorisation request from the acquirer.
	TransactionResponse *CardPaymentTransaction18 `xml:"TxRspn"`
}

func (c *CardPaymentDataSetTransaction5) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction5) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction5) AddEnvironment() *CardPaymentEnvironment16 {
	c.Environment = new(CardPaymentEnvironment16)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction5) AddTransaction() *CardPaymentTransaction12 {
	c.Transaction = new(CardPaymentTransaction12)
	return c.Transaction
}

func (c *CardPaymentDataSetTransaction5) AddTransactionResponse() *CardPaymentTransaction18 {
	c.TransactionResponse = new(CardPaymentTransaction18)
	return c.TransactionResponse
}
