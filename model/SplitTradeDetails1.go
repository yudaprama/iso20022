package model

// Information describing the high level details of a split trade.
type SplitTradeDetails1 struct {

	// Provides information on the status of a foreign exchange trade in the system.
	StatusDetails *TradeData9 `xml:"StsDtls,omitempty"`

	// Amounts of the foreign exchange trade.
	TradeAmounts *AmountsAndValueDate1 `xml:"TradAmts"`

	// Exchange rate as agreed by the traders.
	AgreedRate *AgreedRate1 `xml:"AgrdRate,omitempty"`
}

func (s *SplitTradeDetails1) AddStatusDetails() *TradeData9 {
	s.StatusDetails = new(TradeData9)
	return s.StatusDetails
}

func (s *SplitTradeDetails1) AddTradeAmounts() *AmountsAndValueDate1 {
	s.TradeAmounts = new(AmountsAndValueDate1)
	return s.TradeAmounts
}

func (s *SplitTradeDetails1) AddAgreedRate() *AgreedRate1 {
	s.AgreedRate = new(AgreedRate1)
	return s.AgreedRate
}
