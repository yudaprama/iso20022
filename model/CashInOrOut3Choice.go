package model

// Choice of a payment instrument for the cash-in flow or cash-out flow.
type CashInOrOut3Choice struct {

	// Payment instrument for the cash-in flow.
	CashInPaymentInstrument *PaymentInstrument6Choice `xml:"CshInPmtInstrm"`

	// Payment instrument for the cash-out flow.
	CashOutPaymentInstrument *PaymentInstrument7Choice `xml:"CshOutPmtInstrm"`
}

func (c *CashInOrOut3Choice) AddCashInPaymentInstrument() *PaymentInstrument6Choice {
	c.CashInPaymentInstrument = new(PaymentInstrument6Choice)
	return c.CashInPaymentInstrument
}

func (c *CashInOrOut3Choice) AddCashOutPaymentInstrument() *PaymentInstrument7Choice {
	c.CashOutPaymentInstrument = new(PaymentInstrument7Choice)
	return c.CashOutPaymentInstrument
}
