package model

// Identification of market in which a trade transaction has been executed.
type PlaceOfTradeIdentification1 struct {

	// Identification and type of the place of trade.
	MarketTypeAndIdentification *MarketIdentification84 `xml:"MktTpAndId,omitempty"`

	// Legal entity identification as an alternate identification for a place of trade.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PlaceOfTradeIdentification1) AddMarketTypeAndIdentification() *MarketIdentification84 {
	p.MarketTypeAndIdentification = new(MarketIdentification84)
	return p.MarketTypeAndIdentification
}

func (p *PlaceOfTradeIdentification1) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
