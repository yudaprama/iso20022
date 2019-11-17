package model

// Information about a switch leg that is rejected or repaired.
type SwitchLegReferences1 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	RedemptionLegIdentification *Max35Text `xml:"RedLegId"`

	// Unique technical identifier for an instance of a leg within a switch.
	SubscriptionLegIdentification *Max35Text `xml:"SbcptLegId"`

	// Additional information about the reason for the rejection of a leg.
	LegRejectionReason *Max350Text `xml:"LegRjctnRsn,omitempty"`

	// Elements from the original switch order that have been repaired so that the switch order can be accepted.
	RepairedConditions *RepairedConditions3 `xml:"RprdConds,omitempty"`

	// Account identification of the switch leg that is rejected or repaired.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument identification of the switch leg that is rejected or repaired.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SwitchLegReferences1) SetRedemptionLegIdentification(value string) {
	s.RedemptionLegIdentification = (*Max35Text)(&value)
}

func (s *SwitchLegReferences1) SetSubscriptionLegIdentification(value string) {
	s.SubscriptionLegIdentification = (*Max35Text)(&value)
}

func (s *SwitchLegReferences1) SetLegRejectionReason(value string) {
	s.LegRejectionReason = (*Max350Text)(&value)
}

func (s *SwitchLegReferences1) AddRepairedConditions() *RepairedConditions3 {
	s.RepairedConditions = new(RepairedConditions3)
	return s.RepairedConditions
}

func (s *SwitchLegReferences1) AddInvestmentAccountDetails() *InvestmentAccount13 {
	s.InvestmentAccountDetails = new(InvestmentAccount13)
	return s.InvestmentAccountDetails
}

func (s *SwitchLegReferences1) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}
