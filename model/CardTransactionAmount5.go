package model

// Amounts of the transaction expressed within the terminal currency.
type CardTransactionAmount5 struct {

	// Total amount of the transaction.
	// It corresponds to ISO 8583, field number 4, completed by the field number 49 for the versions 87 and 93.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`

	// Present when cardholder billing currency differs from transaction currency expressed in TransactionAmount. It may be populated by the scheme or intermediary processor as normally Acceptor does not know cardholder billing currency.
	CardholderBillingTransactionAmount *DetailedAmount8 `xml:"CrdhldrBllgTxAmt,omitempty"`

	// Only present within financial transactions when reconciliation currency differs from transaction currency. It may be populated by acquirers in the request or by the schemes in the responses, depending where the reconciliation point is located.
	ReconciliationTransactionAmount *DetailedAmount8 `xml:"RcncltnTxAmt,omitempty"`
}

func (c *CardTransactionAmount5) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewCurrencyAndAmount(value, currency)
}

func (c *CardTransactionAmount5) AddCardholderBillingTransactionAmount() *DetailedAmount8 {
	c.CardholderBillingTransactionAmount = new(DetailedAmount8)
	return c.CardholderBillingTransactionAmount
}

func (c *CardTransactionAmount5) AddReconciliationTransactionAmount() *DetailedAmount8 {
	c.ReconciliationTransactionAmount = new(DetailedAmount8)
	return c.ReconciliationTransactionAmount
}
