package model

// Information about a high frequency trading profile.
type HighFrequencyTradingProfile1 struct {

	// Date on which the investor starts high frequency trading.
	Date *ISODate `xml:"Dt,omitempty"`

	// Frequency of settlement.
	SettlementFrequency *SettlementFrequency1Choice `xml:"SttlmFrqcy,omitempty"`

	// Specifies whether consolidation is done generally or at the level of segregated account.
	ConsolidationType *ConsolidationType1Choice `xml:"CnsldtnTp,omitempty"`
}

func (h *HighFrequencyTradingProfile1) SetDate(value string) {
	h.Date = (*ISODate)(&value)
}

func (h *HighFrequencyTradingProfile1) AddSettlementFrequency() *SettlementFrequency1Choice {
	h.SettlementFrequency = new(SettlementFrequency1Choice)
	return h.SettlementFrequency
}

func (h *HighFrequencyTradingProfile1) AddConsolidationType() *ConsolidationType1Choice {
	h.ConsolidationType = new(ConsolidationType1Choice)
	return h.ConsolidationType
}
