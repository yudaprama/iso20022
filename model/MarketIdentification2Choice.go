package model

// Choice of market identification.
type MarketIdentification2Choice struct {

	// ISO 10383 Market Identification Code.
	MarketIdentifierCode *MICIdentifier `xml:"MktIdrCd"`

	// Description of the market when no Market Identifier Code is available.
	Description *RestrictedFINXMax30Text `xml:"Desc"`
}

func (m *MarketIdentification2Choice) SetMarketIdentifierCode(value string) {
	m.MarketIdentifierCode = (*MICIdentifier)(&value)
}

func (m *MarketIdentification2Choice) SetDescription(value string) {
	m.Description = (*RestrictedFINXMax30Text)(&value)
}
