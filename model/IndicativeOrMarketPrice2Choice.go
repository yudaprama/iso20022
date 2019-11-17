package model

// Choice between an indicative price or a market price.
type IndicativeOrMarketPrice2Choice struct {

	// Estimated price, for example, for valuation purposes.
	IndicativePrice *PriceFormat5Choice `xml:"IndctvPric"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *PriceFormat5Choice `xml:"MktPric"`
}

func (i *IndicativeOrMarketPrice2Choice) AddIndicativePrice() *PriceFormat5Choice {
	i.IndicativePrice = new(PriceFormat5Choice)
	return i.IndicativePrice
}

func (i *IndicativeOrMarketPrice2Choice) AddMarketPrice() *PriceFormat5Choice {
	i.MarketPrice = new(PriceFormat5Choice)
	return i.MarketPrice
}
