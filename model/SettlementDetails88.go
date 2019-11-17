package model

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails88 struct {

	// Indicates the date as known by the two parties to be used for matching purposes when settlement of securities occurs.
	TradeDate *ISODateTime `xml:"TradDt"`

	// Provides details on either the delivering or receiving settlement parties.
	SettlementParties *SettlementParties3Choice `xml:"SttlmPties,omitempty"`

	// Indicates the collateral ownership.
	CollateralOwnership *CollateralOwnership1 `xml:"CollOwnrsh"`
}

func (s *SettlementDetails88) SetTradeDate(value string) {
	s.TradeDate = (*ISODateTime)(&value)
}

func (s *SettlementDetails88) AddSettlementParties() *SettlementParties3Choice {
	s.SettlementParties = new(SettlementParties3Choice)
	return s.SettlementParties
}

func (s *SettlementDetails88) AddCollateralOwnership() *CollateralOwnership1 {
	s.CollateralOwnership = new(CollateralOwnership1)
	return s.CollateralOwnership
}
