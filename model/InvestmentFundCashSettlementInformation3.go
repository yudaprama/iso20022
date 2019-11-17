package model

// Settlement instructions to be used to transfer cash from the Debtor to the Creditor.
type InvestmentFundCashSettlementInformation3 struct {

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument8 `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument9 `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument9 `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument8 `xml:"SvgsPlanPmtInstrm,omitempty"`
}

func (i *InvestmentFundCashSettlementInformation3) AddSubscriptionPaymentInstrument() *PaymentInstrument8 {
	i.SubscriptionPaymentInstrument = new(PaymentInstrument8)
	return i.SubscriptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation3) AddRedemptionPaymentInstrument() *PaymentInstrument9 {
	i.RedemptionPaymentInstrument = new(PaymentInstrument9)
	return i.RedemptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation3) AddDividendPaymentInstrument() *PaymentInstrument9 {
	i.DividendPaymentInstrument = new(PaymentInstrument9)
	return i.DividendPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation3) AddSavingsPlanPaymentInstrument() *PaymentInstrument8 {
	i.SavingsPlanPaymentInstrument = new(PaymentInstrument8)
	return i.SavingsPlanPaymentInstrument
}
