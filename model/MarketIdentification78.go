package model

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification78 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType3Choice `xml:"Tp"`
}

func (m *MarketIdentification78) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

func (m *MarketIdentification78) AddType() *MarketType3Choice {
	m.Type = new(MarketType3Choice)
	return m.Type
}
