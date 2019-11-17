package model

// Card payment completion advice response from the acquirer.
type CardPaymentTransactionAdviceResponse5 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Result of a requested service.
	Response *Response1Code `xml:"Rspn"`
}

func (c *CardPaymentTransactionAdviceResponse5) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionAdviceResponse5) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransactionAdviceResponse5) SetResponse(value string) {
	c.Response = (*Response1Code)(&value)
}
