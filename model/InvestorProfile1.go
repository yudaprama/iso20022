package model

// Information about actions that may be taken on an account.
type InvestorProfile1 struct {

	// Type of profile.
	Type *ProfileType1Choice `xml:"Tp,omitempty"`

	// Status of the profile.
	Status *InvestorProfileStatus1Choice `xml:"Sts,omitempty"`

	// Information about the profile for treasury trading.
	Treasury *TreasuryProfile1 `xml:"Trsr,omitempty"`

	// Information about the profile for high frequency trading.
	HighFrequencyTrading *HighFrequencyTradingProfile1 `xml:"HghFrqcyTradg,omitempty"`

	// Information about the profile for a market marker.
	MarketMaker *MarketMakerProfile1 `xml:"MktMakr,omitempty"`
}

func (i *InvestorProfile1) AddType() *ProfileType1Choice {
	i.Type = new(ProfileType1Choice)
	return i.Type
}

func (i *InvestorProfile1) AddStatus() *InvestorProfileStatus1Choice {
	i.Status = new(InvestorProfileStatus1Choice)
	return i.Status
}

func (i *InvestorProfile1) AddTreasury() *TreasuryProfile1 {
	i.Treasury = new(TreasuryProfile1)
	return i.Treasury
}

func (i *InvestorProfile1) AddHighFrequencyTrading() *HighFrequencyTradingProfile1 {
	i.HighFrequencyTrading = new(HighFrequencyTradingProfile1)
	return i.HighFrequencyTrading
}

func (i *InvestorProfile1) AddMarketMaker() *MarketMakerProfile1 {
	i.MarketMaker = new(MarketMakerProfile1)
	return i.MarketMaker
}
