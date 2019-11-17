package model

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification91 struct {

	// Code allocated to places of trade (stock exchanges), to regulated markets (for example, Electronic Trading Platforms - ECN), and to unregulated markets (for example, Automated Trading Systems - ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification2Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType17Choice `xml:"Tp"`
}

func (m *MarketIdentification91) AddIdentification() *MarketIdentification2Choice {
	m.Identification = new(MarketIdentification2Choice)
	return m.Identification
}

func (m *MarketIdentification91) AddType() *MarketType17Choice {
	m.Type = new(MarketType17Choice)
	return m.Type
}
