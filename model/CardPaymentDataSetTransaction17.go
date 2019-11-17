package model

// Card payment transaction including an authorisation response.
type CardPaymentDataSetTransaction17 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction.
	Environment *CardPaymentEnvironment54 `xml:"Envt"`

	// Card payment transaction authorisation result.
	Transaction *CardPaymentTransaction63 `xml:"Tx"`

	// Response to the authorisation request from the acquirer.
	TransactionResponse *CardPaymentTransaction54 `xml:"TxRspn"`
}

func (c *CardPaymentDataSetTransaction17) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction17) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction17) AddEnvironment() *CardPaymentEnvironment54 {
	c.Environment = new(CardPaymentEnvironment54)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction17) AddTransaction() *CardPaymentTransaction63 {
	c.Transaction = new(CardPaymentTransaction63)
	return c.Transaction
}

func (c *CardPaymentDataSetTransaction17) AddTransactionResponse() *CardPaymentTransaction54 {
	c.TransactionResponse = new(CardPaymentTransaction54)
	return c.TransactionResponse
}
