package model

// Choice of market identification.
type MarketIdentification4Choice struct {

	// Market Identifier Code. Identification of a financial market, as stipulated in the norm ISO 10383 "Codes for exchanges and market identifications".
	MarketIdentifierCode *MICIdentifier `xml:"MktIdrCd"`

	// Description of the market when no Market Identifier Code is available.
	Description *RestrictedFINXMax30Text `xml:"Desc"`
}

func (m *MarketIdentification4Choice) SetMarketIdentifierCode(value string) {
	m.MarketIdentifierCode = (*MICIdentifier)(&value)
}

func (m *MarketIdentification4Choice) SetDescription(value string) {
	m.Description = (*RestrictedFINXMax30Text)(&value)
}
