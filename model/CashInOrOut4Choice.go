package model

// Choice of a payment instrument for the cash-in flow or cash-out flow.
type CashInOrOut4Choice struct {

	// Payment instrument for the cash-in flow.
	CashInPaymentInstrument *PaymentInstrument10Choice `xml:"CshInPmtInstrm"`

	// Payment instrument for the cash-out flow.
	CashOutPaymentInstrument *PaymentInstrument7Choice `xml:"CshOutPmtInstrm"`
}

func (c *CashInOrOut4Choice) AddCashInPaymentInstrument() *PaymentInstrument10Choice {
	c.CashInPaymentInstrument = new(PaymentInstrument10Choice)
	return c.CashInPaymentInstrument
}

func (c *CashInOrOut4Choice) AddCashOutPaymentInstrument() *PaymentInstrument7Choice {
	c.CashOutPaymentInstrument = new(PaymentInstrument7Choice)
	return c.CashOutPaymentInstrument
}
