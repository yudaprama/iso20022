package model

// Amounts of the transaction expressed within the terminal currency.
type CardTransactionAmount3 struct {

	// Total amount of the transaction.
	// It corresponds to ISO 8583, field number 4, completed by the field number 49 for the versions 87 and 93.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount of the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Present when cardholder billing currency differs from transaction currency expressed in TransactionAmount. It may be populated by the scheme or intermediary processor as normally Acceptor does not know cardholder billing currency.
	CardholderBillingTransactionAmount *DetailedAmount8 `xml:"CrdhldrBllgTxAmt,omitempty"`

	// Only present within financial transactions when reconciliation currency differs from transaction currency. It may be populated by acquirers in the request or by the schemes in the responses, depending where the reconciliation point is located.
	ReconciliationTransactionAmount *DetailedAmount8 `xml:"RcncltnTxAmt,omitempty"`

	// Details of the transaction amount, for informational purpose, for instance to be included within cardholder statement.
	// It corresponds partially to ISO 8583, field number 54.
	DetailedAmount []*DetailedAmount9 `xml:"DtldAmt,omitempty"`
}

func (c *CardTransactionAmount3) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewCurrencyAndAmount(value, currency)
}

func (c *CardTransactionAmount3) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardTransactionAmount3) AddCardholderBillingTransactionAmount() *DetailedAmount8 {
	c.CardholderBillingTransactionAmount = new(DetailedAmount8)
	return c.CardholderBillingTransactionAmount
}

func (c *CardTransactionAmount3) AddReconciliationTransactionAmount() *DetailedAmount8 {
	c.ReconciliationTransactionAmount = new(DetailedAmount8)
	return c.ReconciliationTransactionAmount
}

func (c *CardTransactionAmount3) AddDetailedAmount() *DetailedAmount9 {
	newValue := new(DetailedAmount9)
	c.DetailedAmount = append(c.DetailedAmount, newValue)
	return newValue
}
