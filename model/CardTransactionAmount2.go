package model

// Amounts of the transaction expressed within the terminal currency.
type CardTransactionAmount2 struct {

	// Total amount of the transaction.
	// It corresponds to ISO 8583, field number 4, completed by the field number 49 for the versions 87 and 93.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`

	// Present when cardholder billing currency differs from transaction currency expressed in TransactionAmount. It may be populated by the scheme or intermediary processor as normally Acceptor does not know cardholder billing currency.
	CardholderBillingTransactionAmount *DetailedAmount8 `xml:"CrdhldrBllgTxAmt,omitempty"`

	// Details of the TransactionAmount, for informational purposes only, except for cash back which is mandatory for a payment transaction with cashback. The transaction amount is not necessarly the sum of all the detailed amount values.
	DetailedAmount []*DetailedAmount9 `xml:"DtldAmt,omitempty"`
}

func (c *CardTransactionAmount2) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewCurrencyAndAmount(value, currency)
}

func (c *CardTransactionAmount2) AddCardholderBillingTransactionAmount() *DetailedAmount8 {
	c.CardholderBillingTransactionAmount = new(DetailedAmount8)
	return c.CardholderBillingTransactionAmount
}

func (c *CardTransactionAmount2) AddDetailedAmount() *DetailedAmount9 {
	newValue := new(DetailedAmount9)
	c.DetailedAmount = append(c.DetailedAmount, newValue)
	return newValue
}
