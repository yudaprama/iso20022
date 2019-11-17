package model

// Card payment transaction including an authorisation response.
type CardPaymentDataSetTransaction13 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction.
	Environment *CardPaymentEnvironment41 `xml:"Envt"`

	// Card payment transaction authorisation result.
	Transaction *CardPaymentTransaction48 `xml:"Tx"`

	// Response to the authorisation request from the acquirer.
	TransactionResponse *CardPaymentTransaction39 `xml:"TxRspn"`
}

func (c *CardPaymentDataSetTransaction13) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction13) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction13) AddEnvironment() *CardPaymentEnvironment41 {
	c.Environment = new(CardPaymentEnvironment41)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction13) AddTransaction() *CardPaymentTransaction48 {
	c.Transaction = new(CardPaymentTransaction48)
	return c.Transaction
}

func (c *CardPaymentDataSetTransaction13) AddTransactionResponse() *CardPaymentTransaction39 {
	c.TransactionResponse = new(CardPaymentTransaction39)
	return c.TransactionResponse
}
