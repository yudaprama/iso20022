package model

// Details of the transaction in the cancellation response.
type CardPaymentTransactionDetails6 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`
}

func (c *CardPaymentTransactionDetails6) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails6) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}
