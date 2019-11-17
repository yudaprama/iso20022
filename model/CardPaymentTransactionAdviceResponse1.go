package model

// Card payment completion advice response from the acquirer.
type CardPaymentTransactionAdviceResponse1 struct {

	// Unique identification of the transaction by the POI.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Result of a requested service.
	Response *Response1Code `xml:"Rspn"`
}

func (c *CardPaymentTransactionAdviceResponse1) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransactionAdviceResponse1) SetResponse(value string) {
	c.Response = (*Response1Code)(&value)
}
