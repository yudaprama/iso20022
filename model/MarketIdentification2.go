package model

// Provides information about market identification and market type.
type MarketIdentification2 struct {

	// Specifies the type of market.
	Type *MarketTypeFormat1Choice `xml:"Tp"`

	// Identifies the market.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`
}

func (m *MarketIdentification2) AddType() *MarketTypeFormat1Choice {
	m.Type = new(MarketTypeFormat1Choice)
	return m.Type
}

func (m *MarketIdentification2) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}
