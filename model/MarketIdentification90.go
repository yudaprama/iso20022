package model

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification90 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification2Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType16Choice `xml:"Tp"`
}

func (m *MarketIdentification90) AddIdentification() *MarketIdentification2Choice {
	m.Identification = new(MarketIdentification2Choice)
	return m.Identification
}

func (m *MarketIdentification90) AddType() *MarketType16Choice {
	m.Type = new(MarketType16Choice)
	return m.Type
}
