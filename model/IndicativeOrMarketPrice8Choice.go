package model

// Choice between an indicative price or a market price.
type IndicativeOrMarketPrice8Choice struct {

	// Estimated price, for example, for valuation purposes.
	IndicativePrice *PriceFormat50Choice `xml:"IndctvPric"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *PriceFormat50Choice `xml:"MktPric"`
}

func (i *IndicativeOrMarketPrice8Choice) AddIndicativePrice() *PriceFormat50Choice {
	i.IndicativePrice = new(PriceFormat50Choice)
	return i.IndicativePrice
}

func (i *IndicativeOrMarketPrice8Choice) AddMarketPrice() *PriceFormat50Choice {
	i.MarketPrice = new(PriceFormat50Choice)
	return i.MarketPrice
}
