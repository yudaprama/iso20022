package model

// Details of the transaction in the cancellation inside a batch transfer.
type CardPaymentTransactionDetails9 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails9) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails9) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails9) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails9) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
