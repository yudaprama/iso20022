package model

// Choice between an indicative price or a market price.
type IndicativeOrMarketPrice9Choice struct {

	// Estimated price, for example, for valuation purposes.
	IndicativePrice *PriceFormat52Choice `xml:"IndctvPric"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *PriceFormat52Choice `xml:"MktPric"`
}

func (i *IndicativeOrMarketPrice9Choice) AddIndicativePrice() *PriceFormat52Choice {
	i.IndicativePrice = new(PriceFormat52Choice)
	return i.IndicativePrice
}

func (i *IndicativeOrMarketPrice9Choice) AddMarketPrice() *PriceFormat52Choice {
	i.MarketPrice = new(PriceFormat52Choice)
	return i.MarketPrice
}
