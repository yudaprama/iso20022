package model

// Subscription leg, or switch-in, of a switch order.
type SwitchSubscriptionLegOrder6 struct {

	// Unique technical identifier for the instance of the leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Investment fund class related to the subscription leg of the order.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Amount of money, number of units or percentage for the redemption leg of the switch order.
	FinancialInstrumentQuantityChoice *FinancialInstrumentQuantity26Choice `xml:"FinInstrmQtyChc,omitempty"`

	// Account impacted by the investment fund order.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Fees (charges/commission) and tax to be applied to the net amount.
	TransactionOverhead *FeeAndTax1 `xml:"TxOvrhd,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters11 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive/performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SwitchSubscriptionLegOrder6) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegOrder6) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}

func (s *SwitchSubscriptionLegOrder6) AddFinancialInstrumentQuantityChoice() *FinancialInstrumentQuantity26Choice {
	s.FinancialInstrumentQuantityChoice = new(FinancialInstrumentQuantity26Choice)
	return s.FinancialInstrumentQuantityChoice
}

func (s *SwitchSubscriptionLegOrder6) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SwitchSubscriptionLegOrder6) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchSubscriptionLegOrder6) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegOrder6) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegOrder6) AddTransactionOverhead() *FeeAndTax1 {
	s.TransactionOverhead = new(FeeAndTax1)
	return s.TransactionOverhead
}

func (s *SwitchSubscriptionLegOrder6) AddSettlementAndCustodyDetails() *FundSettlementParameters11 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters11)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchSubscriptionLegOrder6) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegOrder6) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SwitchSubscriptionLegOrder6) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SwitchSubscriptionLegOrder6) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}
