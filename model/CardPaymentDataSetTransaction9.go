package model

// Card payment transaction including an authorisation response.
type CardPaymentDataSetTransaction9 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction.
	Environment *CardPaymentEnvironment28 `xml:"Envt"`

	// Card payment transaction authorisation result.
	Transaction *CardPaymentTransaction32 `xml:"Tx"`

	// Response to the authorisation request from the acquirer.
	TransactionResponse *CardPaymentTransaction24 `xml:"TxRspn"`
}

func (c *CardPaymentDataSetTransaction9) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction9) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction9) AddEnvironment() *CardPaymentEnvironment28 {
	c.Environment = new(CardPaymentEnvironment28)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction9) AddTransaction() *CardPaymentTransaction32 {
	c.Transaction = new(CardPaymentTransaction32)
	return c.Transaction
}

func (c *CardPaymentDataSetTransaction9) AddTransactionResponse() *CardPaymentTransaction24 {
	c.TransactionResponse = new(CardPaymentTransaction24)
	return c.TransactionResponse
}
