package model

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification20 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType8Choice `xml:"Tp"`
}

func (m *MarketIdentification20) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

func (m *MarketIdentification20) AddType() *MarketType8Choice {
	m.Type = new(MarketType8Choice)
	return m.Type
}
