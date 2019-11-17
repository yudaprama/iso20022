package model

// Instrument used to process a payment instruction.
type PaymentInstrument13 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Percentage of the dividend payment not to be reinvested, that is, to be paid in cash.
	DividendPercentage *PercentageBoundedRate `xml:"DvddPctg,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument18Choice `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument19Choice `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument19Choice `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument18Choice `xml:"SvgsPlanPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for an interest payment.
	InterestPaymentInstrument *PaymentInstrument19Choice `xml:"IntrstPmtInstrm,omitempty"`
}

func (p *PaymentInstrument13) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument13) SetDividendPercentage(value string) {
	p.DividendPercentage = (*PercentageBoundedRate)(&value)
}

func (p *PaymentInstrument13) AddSubscriptionPaymentInstrument() *PaymentInstrument18Choice {
	p.SubscriptionPaymentInstrument = new(PaymentInstrument18Choice)
	return p.SubscriptionPaymentInstrument
}

func (p *PaymentInstrument13) AddRedemptionPaymentInstrument() *PaymentInstrument19Choice {
	p.RedemptionPaymentInstrument = new(PaymentInstrument19Choice)
	return p.RedemptionPaymentInstrument
}

func (p *PaymentInstrument13) AddDividendPaymentInstrument() *PaymentInstrument19Choice {
	p.DividendPaymentInstrument = new(PaymentInstrument19Choice)
	return p.DividendPaymentInstrument
}

func (p *PaymentInstrument13) AddSavingsPlanPaymentInstrument() *PaymentInstrument18Choice {
	p.SavingsPlanPaymentInstrument = new(PaymentInstrument18Choice)
	return p.SavingsPlanPaymentInstrument
}

func (p *PaymentInstrument13) AddInterestPaymentInstrument() *PaymentInstrument19Choice {
	p.InterestPaymentInstrument = new(PaymentInstrument19Choice)
	return p.InterestPaymentInstrument
}
