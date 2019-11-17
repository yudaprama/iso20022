package model

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails118 struct {

	// Indicates the date as known by the two parties to be used for matching purposes when settlement of securities occurs.
	TradeDate *ISODateTime `xml:"TradDt"`

	// Provides details on either the delivering or receiving settlement parties.
	SettlementParties *SettlementParties7Choice `xml:"SttlmPties,omitempty"`

	// Indicates the collateral ownership.
	CollateralOwnership *CollateralOwnership2 `xml:"CollOwnrsh"`
}

func (s *SettlementDetails118) SetTradeDate(value string) {
	s.TradeDate = (*ISODateTime)(&value)
}

func (s *SettlementDetails118) AddSettlementParties() *SettlementParties7Choice {
	s.SettlementParties = new(SettlementParties7Choice)
	return s.SettlementParties
}

func (s *SettlementDetails118) AddCollateralOwnership() *CollateralOwnership2 {
	s.CollateralOwnership = new(CollateralOwnership2)
	return s.CollateralOwnership
}
