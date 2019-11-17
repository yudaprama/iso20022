package model

// Choice between an indicative price or a market price.
type IndicativeOrMarketPrice11Choice struct {

	// Estimated price, for example, for valuation purposes.
	IndicativePrice *PriceFormat57Choice `xml:"IndctvPric"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *PriceFormat57Choice `xml:"MktPric"`
}

func (i *IndicativeOrMarketPrice11Choice) AddIndicativePrice() *PriceFormat57Choice {
	i.IndicativePrice = new(PriceFormat57Choice)
	return i.IndicativePrice
}

func (i *IndicativeOrMarketPrice11Choice) AddMarketPrice() *PriceFormat57Choice {
	i.MarketPrice = new(PriceFormat57Choice)
	return i.MarketPrice
}
