package model

// Settlement instructions to be used to transfer cash from the debtor to the creditor.
type InvestmentFundCashSettlementInformation8 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument11 `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument12 `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument12 `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument11 `xml:"SvgsPlanPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for an interest payment.
	InterestPaymentInstrument *PaymentInstrument12 `xml:"IntrstPmtInstrm,omitempty"`
}

func (i *InvestmentFundCashSettlementInformation8) SetModificationScopeIndication(value string) {
	i.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (i *InvestmentFundCashSettlementInformation8) AddSubscriptionPaymentInstrument() *PaymentInstrument11 {
	i.SubscriptionPaymentInstrument = new(PaymentInstrument11)
	return i.SubscriptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation8) AddRedemptionPaymentInstrument() *PaymentInstrument12 {
	i.RedemptionPaymentInstrument = new(PaymentInstrument12)
	return i.RedemptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation8) AddDividendPaymentInstrument() *PaymentInstrument12 {
	i.DividendPaymentInstrument = new(PaymentInstrument12)
	return i.DividendPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation8) AddSavingsPlanPaymentInstrument() *PaymentInstrument11 {
	i.SavingsPlanPaymentInstrument = new(PaymentInstrument11)
	return i.SavingsPlanPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation8) AddInterestPaymentInstrument() *PaymentInstrument12 {
	i.InterestPaymentInstrument = new(PaymentInstrument12)
	return i.InterestPaymentInstrument
}
