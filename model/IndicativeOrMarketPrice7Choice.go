package model

// Choice between an indicative price or a market price.
type IndicativeOrMarketPrice7Choice struct {

	// Estimated price, for example, for valuation purposes.
	IndicativePrice *PriceFormat45Choice `xml:"IndctvPric"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *PriceFormat45Choice `xml:"MktPric"`
}

func (i *IndicativeOrMarketPrice7Choice) AddIndicativePrice() *PriceFormat45Choice {
	i.IndicativePrice = new(PriceFormat45Choice)
	return i.IndicativePrice
}

func (i *IndicativeOrMarketPrice7Choice) AddMarketPrice() *PriceFormat45Choice {
	i.MarketPrice = new(PriceFormat45Choice)
	return i.MarketPrice
}
