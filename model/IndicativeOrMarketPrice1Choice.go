package model

// Choice between an indicative price or a market price.
type IndicativeOrMarketPrice1Choice struct {

	// Estimated price, for example, for valuation purposes.
	IndicativePrice *PriceFormat11Choice `xml:"IndctvPric"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *PriceFormat11Choice `xml:"MktPric"`
}

func (i *IndicativeOrMarketPrice1Choice) AddIndicativePrice() *PriceFormat11Choice {
	i.IndicativePrice = new(PriceFormat11Choice)
	return i.IndicativePrice
}

func (i *IndicativeOrMarketPrice1Choice) AddMarketPrice() *PriceFormat11Choice {
	i.MarketPrice = new(PriceFormat11Choice)
	return i.MarketPrice
}
