package model

// Specifies prices.
type CorporateActionPrice4 struct {

	// Estimated price, eg, for valuation purposes.
	IndicativePrice *PriceFormat2Choice `xml:"IndctvPric,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *PriceFormat2Choice `xml:"MktPric,omitempty"`
}

func (c *CorporateActionPrice4) AddIndicativePrice() *PriceFormat2Choice {
	c.IndicativePrice = new(PriceFormat2Choice)
	return c.IndicativePrice
}

func (c *CorporateActionPrice4) AddMarketPrice() *PriceFormat2Choice {
	c.MarketPrice = new(PriceFormat2Choice)
	return c.MarketPrice
}
