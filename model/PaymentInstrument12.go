package model

// Instrument that has or represents monetary value and is used to process a payment instruction.
type PaymentInstrument12 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Choice of payment instruments.
	PaymentInstrument *PaymentInstrument16Choice `xml:"PmtInstrm"`

	// Percentage of the dividend payment not to be reinvested, that is, to be paid in cash.
	DividendPercentage *PercentageBoundedRate `xml:"DvddPctg,omitempty"`
}

func (p *PaymentInstrument12) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument12) AddPaymentInstrument() *PaymentInstrument16Choice {
	p.PaymentInstrument = new(PaymentInstrument16Choice)
	return p.PaymentInstrument
}

func (p *PaymentInstrument12) SetDividendPercentage(value string) {
	p.DividendPercentage = (*PercentageBoundedRate)(&value)
}
