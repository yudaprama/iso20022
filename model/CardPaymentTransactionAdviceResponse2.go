package model

// Cancellation advice response from the acquirer.
type CardPaymentTransactionAdviceResponse2 struct {

	// Unique identification of the transaction by the POI.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`
}

func (c *CardPaymentTransactionAdviceResponse2) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}
