package model

// Transaction information in the cancellation request.
type CardPaymentTransaction5 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction8 `xml:"OrgnlTx"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails5 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction5) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction5) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction5) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction5) AddOriginalTransaction() *CardPaymentTransaction8 {
	c.OriginalTransaction = new(CardPaymentTransaction8)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction5) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction5) AddTransactionDetails() *CardPaymentTransactionDetails5 {
	c.TransactionDetails = new(CardPaymentTransactionDetails5)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction5) SetAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = (*Max70Text)(&value)
}
