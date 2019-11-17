package model

// Information describing the high level details of a split trade.
type SplitTradeDetails3 struct {

	// Provides information on the status of a foreign exchange trade in the system.
	StatusDetails *TradeData16 `xml:"StsDtls,omitempty"`

	// Amounts of the foreign exchange trade.
	TradeAmounts *AmountsAndValueDate1 `xml:"TradAmts"`

	// Exchange rate as agreed by the traders.
	AgreedRate *AgreedRate3 `xml:"AgrdRate,omitempty"`
}

func (s *SplitTradeDetails3) AddStatusDetails() *TradeData16 {
	s.StatusDetails = new(TradeData16)
	return s.StatusDetails
}

func (s *SplitTradeDetails3) AddTradeAmounts() *AmountsAndValueDate1 {
	s.TradeAmounts = new(AmountsAndValueDate1)
	return s.TradeAmounts
}

func (s *SplitTradeDetails3) AddAgreedRate() *AgreedRate3 {
	s.AgreedRate = new(AgreedRate3)
	return s.AgreedRate
}
