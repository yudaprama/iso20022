package model

// Instrument that has or represents monetary value and is used to process a payment instruction.
type PaymentInstrument10 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Choice of payment instruments.
	PaymentInstrument *PaymentInstrument16Choice `xml:"PmtInstrm"`

	// Percentage of the dividend payment not to be reinvested.
	DividendPercentage *PercentageBoundedRate `xml:"DvddPctg,omitempty"`
}

func (p *PaymentInstrument10) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument10) AddPaymentInstrument() *PaymentInstrument16Choice {
	p.PaymentInstrument = new(PaymentInstrument16Choice)
	return p.PaymentInstrument
}

func (p *PaymentInstrument10) SetDividendPercentage(value string) {
	p.DividendPercentage = (*PercentageBoundedRate)(&value)
}
