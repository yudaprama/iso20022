package model

// Choice of market identification.
type MarketIdentification1Choice struct {

	// ISO 10383 Market Identification Code.
	MarketIdentifierCode *MICIdentifier `xml:"MktIdrCd"`

	// Description of the market when no Market Identifier Code is available.
	Description *Max35Text `xml:"Desc"`
}

func (m *MarketIdentification1Choice) SetMarketIdentifierCode(value string) {
	m.MarketIdentifierCode = (*MICIdentifier)(&value)
}

func (m *MarketIdentification1Choice) SetDescription(value string) {
	m.Description = (*Max35Text)(&value)
}
