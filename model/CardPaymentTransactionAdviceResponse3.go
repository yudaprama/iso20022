package model

// Card payment completion advice response from the acquirer.
type CardPaymentTransactionAdviceResponse3 struct {

	// Unique identification of the transaction by the POI.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Result of a requested service.
	Response *Response1Code `xml:"Rspn"`
}

func (c *CardPaymentTransactionAdviceResponse3) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransactionAdviceResponse3) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionAdviceResponse3) SetResponse(value string) {
	c.Response = (*Response1Code)(&value)
}
