package model

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification13 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification3Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType12Choice `xml:"Tp,omitempty"`
}

func (m *MarketIdentification13) AddIdentification() *MarketIdentification3Choice {
	m.Identification = new(MarketIdentification3Choice)
	return m.Identification
}

func (m *MarketIdentification13) AddType() *MarketType12Choice {
	m.Type = new(MarketType12Choice)
	return m.Type
}
