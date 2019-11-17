package model

// Settlement instructions to be used to transfer cash from the Debtor to the Creditor.
type InvestmentFundCashSettlementInformation6 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument11 `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument10 `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument10 `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument11 `xml:"SvgsPlanPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for an interest payment.
	InterestPaymentInstrument *PaymentInstrument10 `xml:"IntrstPmtInstrm,omitempty"`
}

func (i *InvestmentFundCashSettlementInformation6) SetModificationScopeIndication(value string) {
	i.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (i *InvestmentFundCashSettlementInformation6) AddSubscriptionPaymentInstrument() *PaymentInstrument11 {
	i.SubscriptionPaymentInstrument = new(PaymentInstrument11)
	return i.SubscriptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation6) AddRedemptionPaymentInstrument() *PaymentInstrument10 {
	i.RedemptionPaymentInstrument = new(PaymentInstrument10)
	return i.RedemptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation6) AddDividendPaymentInstrument() *PaymentInstrument10 {
	i.DividendPaymentInstrument = new(PaymentInstrument10)
	return i.DividendPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation6) AddSavingsPlanPaymentInstrument() *PaymentInstrument11 {
	i.SavingsPlanPaymentInstrument = new(PaymentInstrument11)
	return i.SavingsPlanPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation6) AddInterestPaymentInstrument() *PaymentInstrument10 {
	i.InterestPaymentInstrument = new(PaymentInstrument10)
	return i.InterestPaymentInstrument
}
