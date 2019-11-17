package model

// Choice of a payment instrument for the cash-in flow or cash-out flow.
type CashInOrOut7Choice struct {

	// Payment instrument for the cash-in flow.
	CashInPaymentInstrument *PaymentInstrument20Choice `xml:"CshInPmtInstrm"`

	// Payment instrument for the cash-out flow.
	CashOutPaymentInstrument *PaymentInstrument21Choice `xml:"CshOutPmtInstrm"`
}

func (c *CashInOrOut7Choice) AddCashInPaymentInstrument() *PaymentInstrument20Choice {
	c.CashInPaymentInstrument = new(PaymentInstrument20Choice)
	return c.CashInPaymentInstrument
}

func (c *CashInOrOut7Choice) AddCashOutPaymentInstrument() *PaymentInstrument21Choice {
	c.CashOutPaymentInstrument = new(PaymentInstrument21Choice)
	return c.CashOutPaymentInstrument
}
