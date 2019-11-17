package model

// Instrument that has or represents monetary value and is used to process a payment instruction.
type PaymentInstrument11 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Choice of payment instruments.
	PaymentInstrument *PaymentInstrument17Choice `xml:"PmtInstrm"`
}

func (p *PaymentInstrument11) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument11) AddPaymentInstrument() *PaymentInstrument17Choice {
	p.PaymentInstrument = new(PaymentInstrument17Choice)
	return p.PaymentInstrument
}
