package model

// Choice between an indicative price or a market price.
type IndicativeOrMarketPrice5Choice struct {

	// Estimated price, for example, for valuation purposes.
	IndicativePrice *PriceFormat19Choice `xml:"IndctvPric"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *PriceFormat19Choice `xml:"MktPric"`
}

func (i *IndicativeOrMarketPrice5Choice) AddIndicativePrice() *PriceFormat19Choice {
	i.IndicativePrice = new(PriceFormat19Choice)
	return i.IndicativePrice
}

func (i *IndicativeOrMarketPrice5Choice) AddMarketPrice() *PriceFormat19Choice {
	i.MarketPrice = new(PriceFormat19Choice)
	return i.MarketPrice
}
