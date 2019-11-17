package model

// Choice of a payment instrument for the cash-in flow or cash-out flow.
type CashInOrOut5Choice struct {

	// Payment instrument for the cash-in flow.
	CashInPaymentInstrument *PaymentInstrument12Choice `xml:"CshInPmtInstrm"`

	// Payment instrument for the cash-out flow.
	CashOutPaymentInstrument *PaymentInstrument11Choice `xml:"CshOutPmtInstrm"`
}

func (c *CashInOrOut5Choice) AddCashInPaymentInstrument() *PaymentInstrument12Choice {
	c.CashInPaymentInstrument = new(PaymentInstrument12Choice)
	return c.CashInPaymentInstrument
}

func (c *CashInOrOut5Choice) AddCashOutPaymentInstrument() *PaymentInstrument11Choice {
	c.CashOutPaymentInstrument = new(PaymentInstrument11Choice)
	return c.CashOutPaymentInstrument
}
