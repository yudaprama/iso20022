package model

// Indicates whether the trade was executed on or off-market.
type PlaceOfTradeIdentification2Choice struct {

	// Exchange or Multilateral Trading Facility (MTF) on which the transaction is executed.
	MarketIdentification *MICIdentifier `xml:"MktId"`

	// Indicates that the trade was executed off -exchange.
	OffMarket *OffMarket1Choice `xml:"OffMkt"`
}

func (p *PlaceOfTradeIdentification2Choice) SetMarketIdentification(value string) {
	p.MarketIdentification = (*MICIdentifier)(&value)
}

func (p *PlaceOfTradeIdentification2Choice) AddOffMarket() *OffMarket1Choice {
	p.OffMarket = new(OffMarket1Choice)
	return p.OffMarket
}
